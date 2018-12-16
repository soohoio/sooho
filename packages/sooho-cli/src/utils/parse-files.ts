import * as crypto from 'crypto'
import * as fs from 'fs'
import {promisify} from 'util'
import * as parser from '@sooho/parser'

export async function parseFiles(filePaths, abstract) {
  let functions = []
  let constructors = []
  let parseErrors = []
  let totalLines = 0

  const readFile = promisify(fs.readFile)

  await Promise.all(filePaths.map(async filePath => {
    const input = await readFile(filePath, 'utf8')
    totalLines += input.toString().split('\n').length - 1
    try {
      if (!input) throw new Error('Invalid input')
      const ast = parser.parse(input, {loc: true, abstract})
      parser.visit(ast, {
        FunctionDefinition: node => {
          const body = node.self.getText()
          const result = {
            filePath,
            loc: {
              startLine: node.loc.start.line,
              endLine: node.loc.end.line
            },
            signature: crypto.createHash('md5').update(body).digest('hex')
          }
          if (node.isConstructor) {
            constructors.push(result)
          } else {
            functions.push(result)
          }
        }
      })
    } catch (e) {
      if (e instanceof parser.ParserError) {
        if (process.env.DEBUG) {
          console.error(e.errors)
        }
        parseErrors.push({
          filePath,
          messages: e.errors
        })
      }
    }
  }))

  return {
    version: '0.3.0',
    totalFiles: filePaths.length,
    totalLines,
    totalSigns: functions.length + constructors.length,
    success: {
      functions,
      constructors
    },
    errors: parseErrors
  }
}
