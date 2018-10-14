import {Command, flags} from '@oclif/command'
import * as parser from '@sooho/parser'
import * as fs from 'fs'
import {promisify} from 'util'

export default class Encrypt extends Command {
  static description = 'Encrypt source code into hash file'

  static examples = [
    '$ sooho encrypt FILE_PATH -o result',
  ]

  static flags = {
    help: flags.help({char: 'h', description: 'show CLI help'}),
    // flag with a value (-o, --output=VALUE)
    output: flags.string({char: 'o', description: 'name to save'}),
    // flag with no value (-f, --force)
    force: flags.boolean({char: 'f'}),
  }

  static args = [{name: 'filePath', required: true, description: 'entry path'}]

  async run() {
    const {args, flags} = this.parse(Encrypt)

    const output = flags.output || `${args.filePath}.aegis`
    this.log(`Extracting ${args.filePath} into ${output} from ${__filename}`)

    const readFile = promisify(fs.readFile)
    const input = await readFile(args.filePath, 'utf8')

    try {
      if (!input) throw new Error('Invalid input')
      const ast = parser.parse(input, {loc: true})
      const functions = []
      parser.visit(ast, {
        FunctionDefinition: node => {
          if (node.isConstructor) {
            // console.warn("[Constructor]")
            // console.warn(`Code: ${node.ctx.getText()}`)
          } else {
            this.log('[Function]')
            this.log(node.ctx.getText())
          }
        }
      })
    } catch (e) {
      if (e instanceof parser.ParserError) {
        this.log(e.errors)
      }
    }
  }
}
