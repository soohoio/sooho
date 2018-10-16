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
    abstract: flags.boolean({char: 'a', description: 'turn on abstraction mode'}),
    help: flags.help({char: 'h', description: 'show CLI help'}),
    save: flags.boolean({char: 's', description: 'save encrypted result'}),
  }

  static args = [{name: 'filePath', required: true, description: 'entry path'}]

  async run() {
    const {args, flags} = this.parse(Encrypt)
    const filePath = args.filePath
    const abstract = flags.abstract || false
    const save = flags.save || false

    this.log(
      `Extracting signatures in ${filePath}`,
      `${abstract ? 'with' : 'without'} abstraction`
    )

    const readFile = promisify(fs.readFile)
    const input = await readFile(filePath, 'utf8')
    const functions = []

    try {
      if (!input) throw new Error('Invalid input')
      const ast = parser.parse(input, {loc: true, abstract})
      parser.visit(ast, {
        FunctionDefinition: node => {
          if (node.isConstructor) {
            // console.warn("[Constructor]")
            // console.warn(`Code: ${node.ctx.getText()}`)
          } else {
            const body = node.self.getText()
            functions.push({
              loc: JSON.stringify({
                startLine: node.loc.start.line,
                endLine: node.loc.end.line
              }),
              signature: crypto.createHash('md5').update(body).digest('hex'),
            })
          }
        }
      })
    } catch (e) {
      if (e instanceof parser.ParserError) {
        this.log(e.errors)
        return
      }
    }

    if (save) {
      const output = fs.createWriteStream(`${filePath.split('.sol')[0]}.aegis`, 'utf8')
      functions.forEach(func => {
        output.write(`[Function]\n${func.loc}\n${func.signature}\n`)
      })
      output.end()
    } else {
      functions.forEach(func => {
        this.log(`[Function]\n${func.loc}\n${func.signature}`)
      })
    }
  }
}
