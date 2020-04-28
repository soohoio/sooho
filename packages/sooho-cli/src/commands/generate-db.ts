import {Command} from '@oclif/command'
import {parseFiles} from '../utils/parse-files'
import {promisify} from 'util'
import * as path from 'path'
import * as fs from 'fs'
import * as Diff from 'diff'
import * as YAML from 'yaml'
import * as ora from 'ora'

export default class GenerateDB extends Command {
  static description = 'Generate advisory db'

  static examples = [
    '$ sooho generate ORIGIN_SOL_FILE_PATH PATCHED_SOL_FILE_PATH OUTPUT_FILE_PATH',
  ]

  static args = [{
    name: 'originFilePath',
    required: true,
    description: 'Original vulnerable solidity file path'
  }, {
    name: 'patchedFilePath',
    required: true,
    description: 'Patched safe solidity file path'
  }, {
    name: 'outputFilePath',
    required: true,
    description: 'File path to generate yaml file will be located at'
  }]

  async run() {
    const readFile = promisify(fs.readFile)
    const writeFile = promisify(fs.writeFile)

    const {args: {originFilePath, patchedFilePath, outputFilePath}} = this.parse(GenerateDB)
    const spinner = ora({spinner: 'dots'})
    const originFileName = path.basename(originFilePath, '.sol')

    spinner.start('Parse files')
    const parseResult = await parseFiles([originFilePath, patchedFilePath])
    if (parseResult.errors.length) {
      spinner.fail('Parse files')
    } else {
      spinner.succeed('Parse files')
    }

    const originFile = await readFile(originFilePath)
    const patchFile = await readFile(patchedFilePath)

    const functions = parseResult.success.functions
    const funcsInOrigin = functions.filter(f => f.filePathIdx === 0)
    const funcsInPatch = functions.filter(f => f.filePathIdx === 1)

    // filter patched function only
    const patchedFuncs = funcsInPatch.map(f => {
      const { signature, contractName } = f
      const matchedFunc = funcsInOrigin.find(fo => fo.contractName === contractName && fo.signature === signature)
      if (matchedFunc) {
        return null
      }
      return f.functionName
    }).filter(f => f)

    spinner.start('Generate advisory db')
    const generateYaml = await patchedFuncs.map(async funcName => {
      // get original function body
      const originFunc = funcsInOrigin.find(f => f.functionName === funcName)
      const { startLine: originStart, endLine: originEnd } = originFunc.loc
      const originFileBody = originFile.toString()
      const originFuncBody = originFileBody.split('\n').slice(originStart-1, originEnd).join('\n')

      // get patched function body
      const patchFunc = funcsInPatch.find(f => f.functionName === funcName)
      const { startLine: patchStart, endLine: patchEnd } = patchFunc.loc
      const patchFileBody = patchFile.toString()
      const patchFuncBody = patchFileBody.split('\n').slice(patchStart-1, patchEnd).join('\n')

      // get vulnerable function's signature
      const { signature } = originFunc

      // generate patch information
      const diff = Diff.diffLines(originFuncBody, patchFuncBody)
      let patch = ''
      diff.forEach(d => {
        if (d.added || d.removed) {
          const dSlice = d.value.split('\n')
          dSlice.forEach(ds => patch += d.added ? `+${ds}\n` : `-${ds}\n`)
        } else {
          patch += d.value
        }
      })

      // yaml format with dummy infos, signature & patch info
      const yaml = {
        id: '',
        title: '',
        description: '',
        references: ['', ''],
        credits: '',
        vulnerability_type: {
          cwe: '',
          swc: '',
        },
        severity: 0.0,
        affected: {
          contractName: '',
          address: '',
        },
        signature,
        patch,
      }

      // write yaml file into disk
      await writeFile(`${outputFilePath}/${originFileName}-${funcName}.yml`, YAML.stringify(yaml))
      this.log(`\nadvisory-db for function ${funcName} is generated at ${outputFilePath}/${originFileName}-${funcName}.yml`)
    })

    await Promise.all(generateYaml)
    spinner.succeed('Generate advisory db')
  }
}
