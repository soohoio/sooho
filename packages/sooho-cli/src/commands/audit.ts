import {Command} from '@oclif/command'
import {getAdvisoryDB} from '@sooho/advisory-db'
import * as fs from 'fs'
import * as ora from 'ora'
import * as path from 'path'
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
    const db = getAdvisoryDB()

    spinner.start('Checking vulnerabilities')
    functions.forEach(func => {
      if (db.signature == func.signature) {
        isSafe = false
        vulns.push(db)
      }
    })

    if (isSafe) {
      spinner.succeed('Done!')
    } else {
      const table = new Table({
        chars: { 'top': '═' , 'top-mid': '╤' , 'top-left': '╔' , 'top-right': '╗'
         , 'bottom': '═' , 'bottom-mid': '╧' , 'bottom-left': '╚' , 'bottom-right': '╝'
         , 'left': '║' , 'left-mid': '╟' , 'mid': '─' , 'mid-mid': '┼'
         , 'right': '║' , 'right-mid': '╢' , 'middle': '│' },
        head: ['CVE ID', 'Type', 'Severity', 'Desc'],
        colWidths: [20, 10, 10, 50]
      })

      table.push(
        [db.id, db.vulnerability_type.swc, db.severity, db.description]
      )

      spinner.fail(`Vulnerabilities have detected!\n${table.toString()}`)
    }
  }
}
