import * as crypto from 'crypto'
import * as fs from 'fs'
import {promisify} from 'util'
import * as parser from '@sooho/parser'

export async function parseFiles(filePaths) {
  const totalFiles = filePaths.length
  let files = new Array(totalFiles)
  let functions = []
  let constructors = []
  let parseErrors = []
  let totalLines = 0

  const readFile = promisify(fs.readFile)

  await Promise.all(filePaths.map(async (filePath, index) => {
    const input = await readFile(filePath, 'utf8')
    files[index] = {filePath, lines: input.split('\n').length - 1}
    totalLines += files[index].lines
    try {
      if (!input) throw new Error('Invalid input')
      const ast = parser.parse(input, {loc: true})
      let funcId = 0
      let sourceInstance = {}

      const findContractCtx = ctx => {
        const { parentCtx } = ctx

        if (parentCtx.ruleIndex === 9) {
          return parentCtx
        }

        return findContractCtx(parentCtx)
      }

      const getContractName = node => {
        const contractCtx = findContractCtx(node.self)

        return contractCtx.getChild(1).getChild(0).symbol.text
      }

      // find function definition context recursively
      const findFuncDefCtx = ctx => {
        const { parentCtx } = ctx

        // SolidityParser.RULE_functionDefinition = 18
        if (parentCtx.ruleIndex === 18) {
          return parentCtx
        }
        // SolidityParser.RULE_modifierDefinition = 16
        // dont care about modifier
        if (parentCtx.ruleIndex === 16) {
          return null
        }
        // SolidityParser.RULE_contractPart = 11
        // if there was no ruleIndex 18 before 11, this variable is global
        if (parentCtx.ruleIndex === 11) {
          return null
        }

        return findFuncDefCtx(parentCtx)
      }

      const getFunctionName = node => {
        const funcDefCtx = findFuncDefCtx(node.self)
        if (!funcDefCtx) {
          return ''
        }

        let name = funcDefCtx.getChild(1).getChild(0).symbol.text
        if (name === '(') {
          name = 'fallback'
        }

        return name
      }

      const getVariableName = node => {
        return node.name
      }

      const getDataType = node => {
        return node.typeName.name
      }

      const getFunctionCallName = node => {
        let funcCallName = node.expression.name
        // in case of MemberAccess, ElementaryTypeNameExpression
        if (!funcCallName) {
          switch (node.expression.type) {
            case 'MemberAccess':
              return node.expression.memberName
            // ElementaryTypeNameExpression will be handled on data type
            // ex> address(0) -> DTYPE(0), not FUNCCALL(0)
            case 'ElementaryTypeNameExpression':
            default:
              return undefined
          }
        }

        return funcCallName
      }

      parser.visit(ast, {
        FunctionCall: node => {
          const contractName = getContractName(node)
          const functionName = getFunctionName(node)
          if (functionName === '') {
            return
          }

          const functionCallName = getFunctionCallName(node)
          if (functionCallName == undefined) {
            return
          }

          sourceInstance[contractName][functionName].functionCalls.push(functionCallName)
        },
        ContractDefinition: node => {
          const contractName = node.self.getChild(1).getChild(0).symbol.text
          sourceInstance[contractName] = {
            default: {
              dataTypes: [],
              localVariables: []
            }
          }
        },
        VariableDeclaration: node => {
          const contractName = getContractName(node)
          let functionName = getFunctionName(node)
          if (functionName === '') {
            functionName = 'default'
          }

          const variableName = getVariableName(node)
          sourceInstance[contractName][functionName].localVariables.push(variableName)
          const dataType = getDataType(node)
          sourceInstance[contractName][functionName].dataTypes.push(dataType)
        },
        FunctionDefinition: node => {
          const {
            name,
            parameters: { parameters },
            isConstructor,
            loc
          } = node

          const paramNames = parameters.map(param => param.name)
          const contractName = getContractName(node)

          const { start, end } = node.loc
          const { line: startLine } = start
          const { line: endLine } = end

          sourceInstance[contractName][name ? name : 'fallback'] = {
            body: node.self.getText(),
            isConstructor, 
            loc: { startLine, endLine },
            dataTypes: [],
            localVariables: [],
            parameters: paramNames,
            functionCalls: []
          }
        }
      })

      const abstractFunction = (body, functionInstance, globalInstance) => {
        const { parameters, dataTypes, localVariables, functionCalls } = functionInstance
        const { localVariables: globalVariables, dataTypes: globalDataTypes } = globalInstance

        // level 1
        parameters.forEach(param => {
          const pattern = new RegExp("(^|\\W)" + param + "(\\W)", 'gi')
          body = body.replace(pattern, '$1FPARAM$2')
        })

        // level 2
        dataTypes.forEach(dtype => {
          const pattern = new RegExp("(^|\\W)" + dtype + "(\\W)", 'gi')
          body = body.replace(pattern, '$1DTYPE$2')
        })

        globalDataTypes.forEach(dtype => {
          const pattern = new RegExp("(^|\\W)" + dtype + "(\\W)", 'gi')
          body = body.replace(pattern, '$1DTYPE$2')
        })

        // level 3
        localVariables.forEach(lvar => {
          const pattern = new RegExp("(^|\\W)" + lvar + "(\\W)", 'gi')
          body = body.replace(pattern, '$1LVAR$2')
        })

        globalVariables.push('now')
        globalVariables.forEach(gvar => {
          const pattern = new RegExp("(^|\\W)" + gvar + "(\\W)", 'gi')
          body = body.replace(pattern, '$1GVAR$2')
        })

        // level 4
        functionCalls.forEach(fcall => {
          const pattern = new RegExp("(^|\\W)" + fcall + "(\\W)", 'gi')
          body = body.replace(pattern, '$1FUNCCALL$2')
        })

        return body
      }

      const normalizeFunction = (originalBody) => {
        // remove left curly brace & dump function identifier line
        let normBody = originalBody.split('{').slice(1).join('\n')
        // remove right curly brace & get only function body
        normBody = normBody.split('}').slice(0, -1).join('\n')
        // remove spaces
        normBody = normBody.replace(/\n/g, '').replace(/ /g, '').replace(/\t/g, '').replace(/\r/g, '')
        normBody = normBody.toLowerCase()
        return normBody
      }

      Object.keys(sourceInstance).forEach(contract => {
        const contractInstance = sourceInstance[contract]

        for (let functionName in contractInstance) {
          if (functionName === 'default') {
            continue
          }

          const functionInstance = contractInstance[functionName]
          const globalInstance = contractInstance['default']
          const { body } = functionInstance
          const absBody = abstractFunction(body, functionInstance, globalInstance)

          const normAbsBody = normalizeFunction(absBody)
          const { startLine, endLine } = functionInstance.loc

          const result = {
            filePathIdx: index,
            functionName,
            contractName: contract,
            loc: { startLine, endLine },
            signature: crypto.createHash('md5').update(normAbsBody).digest('hex')
          }

          if (functionInstance.isConstructor) {
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
    version: '0.3.4',
    fileInfo: {
      totalFiles,
      totalLines,
      totalSigns: functions.length + constructors.length,
      files
    },
    success: {
      functions,
      constructors
    },
    errors: parseErrors
  }
}
