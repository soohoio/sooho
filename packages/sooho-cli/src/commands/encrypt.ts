import {Command, flags} from '@oclif/command'
import * as parser from '@sooho/parser'
import * as crypto from 'crypto'
import * as fs from 'fs'
import {promisify} from 'util'

export default class Encrypt extends Command {
  static description = 'Encrypt source code into hash file'

  static examples = [
    '$ sooho encrypt FILE_PATH',
  ]

  static flags = {
    help: flags.help({char: 'h', description: 'show CLI help'}),
    abstract: flags.boolean({char: 'a', description: 'turn on abstraction mode'}),
  }

  static args = [{name: 'filePath', required: true, description: 'entry path'}]

  async run() {
    const {args, flags} = this.parse(Encrypt)
    const filePath = args.filePath;
    const abstract = flags.abstract || false;

    this.log(
      `Extracting signatures in ${filePath}`,
      `${abstract ? 'with' : 'without'} abstraction`
    )

    const readFile = promisify(fs.readFile)
    const input = await readFile(filePath, 'utf8')

    try {
      if (!input) throw new Error('Invalid input')
      const ast = parser.parse(input, {loc: true, abstract})
      parser.visit(ast, {
        FunctionDefinition: node => {
          if (node.isConstructor) {
            // console.warn("[Constructor]")
            // console.warn(`Code: ${node.ctx.getText()}`)
          } else {
            this.log('[Function]')
            const body = node.self.getText()
            this.log(JSON.stringify(node.loc))
            this.log(crypto.createHash('md5').update(body).digest("hex"))
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
