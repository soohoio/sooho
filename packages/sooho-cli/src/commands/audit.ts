import {Command} from '@oclif/command'
import {getAdvisoryDB} from '@sooho/advisory-db'
import * as fs from 'fs'
import * as ora from 'ora'
import * as powerwalker from 'powerwalker'
import * as Table from 'cli-table'
import {promisify} from 'util'
import {onlySolidity} from '../utils/filters'
import {parseFiles} from '../utils/parse-files'

export default class Audit extends Command {
  static description = 'Audit smart contract'

  static examples = [
    '$ sooho audit INPUT_PATH',
  ]

  static args = [{name: 'inputPath', required: true, description: 'entry path'}]

  async run() {
    const {args: {inputPath}} = this.parse(Audit)

    const spinner = ora({text: 'Parse files', spinner: 'dots'}).start()
    const lstat = promisify(fs.lstat)
    const stats = await lstat(inputPath)
    const routes = stats.isFile() ? [inputPath] : await powerwalker(inputPath)
    const filePaths = routes.filter(onlySolidity)
    const parsed = await parseFiles(filePaths, true)
    const {errors, success: {functions, constructors}} = parsed

    if (errors.length > 0) {
      if (functions.length > 0 || constructors.length > 0) {
        spinner.warn('Parse files')
      } else {
        spinner.fail('Parse files')
      }
    } else {
      spinner.succeed('Parse files')
    }

    let isSafe = true
    let vulns = []
    try {
      const signatures = functions.map(func => func.signature)
      const db = getAdvisoryDB()

      spinner.start('Checking vulnerabilities')
      db.forEach(cve => {
        if (signatures.indexOf(cve.signature) > 0) {
          isSafe = false
          vulns.push(cve)
        }
      })

      if (isSafe) {
        spinner.succeed('Done!')
      } else {
        const table = new Table({
          head: ['CVE ID', 'Type', 'Severity', 'Desc'],
          colWidths: [20, 10, 10, 50]
        })

        vulns.map(vul => table.push([
          vul.id,
          vul.vulnerability_type.swc,
          vul.severity,
          vul.description
        ]))

        spinner.fail('Vulnerabilities have detected!')
        this.log(table.toString())
      }
    } catch (e) {
      spinner.fail(e)
    }
  }
}
