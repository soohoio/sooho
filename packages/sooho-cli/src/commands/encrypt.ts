import {Command, flags} from '@oclif/command'
import * as parser from '@sooho/parser'
import * as fs from 'fs'
import {promisify} from 'util'

export default class Encrypt extends Command {
  static description = 'Encrypt source code into hash file'

  static examples = [
    '$ sooho encrypt FILE_PATH',
  ]

  static flags = {
    help: flags.help({char: 'h', description: 'show CLI help'}),
    // flag with no value (-f, --force)
    force: flags.boolean({char: 'f'}),
  }

  static args = [{name: 'filePath', required: true, description: 'entry path'}]

  async run() {
    const {args, flags} = this.parse(Encrypt)

    this.log(`Extracting functions from ${args.filePath}`)

    const readFile = promisify(fs.readFile)
    const input = await readFile(args.filePath, 'utf8')

    try {
      if (!input) throw new Error('Invalid input')
      const ast = parser.parse(input, {loc: true})
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
