import {Command} from '@oclif/command'
import * as fs from 'fs'
import * as ora from 'ora'
import * as powerwalker from 'powerwalker'
import {promisify} from 'util'
import {onlySolidity} from '../utils/filters'
import {abstract, help, save} from '../utils/flags'
import {parseFiles} from '../utils/parse-files'

export default class Encrypt extends Command {
  static description = 'Encrypt source code into hash file'

  static examples = [
    '$ sooho encrypt PATH',
  ]

  static flags = {abstract, help, save}
  static args = [{name: 'path', required: true, description: 'entry path'}]

  async run() {
    const {args: {path}, flags: {abstract, save}} = this.parse(Encrypt)

    const spinner = ora({text: 'Parse files', spinner: 'dots'}).start()
    const lstat = promisify(fs.lstat)
    const stats = await lstat(path)
    const routes = stats.isFile() ? [path] : await powerwalker(path)
    const filePaths = routes.filter(onlySolidity)
    const parsed = await parseFiles(filePaths, abstract)
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

    const result = JSON.stringify(parsed, null, 4)

    if (save) {
      spinner.start('Saving results')
      const fileName = `${path.split('/').pop().split('.sol')[0]}.aegis`
      fs.writeFile(fileName, result, err => {
        if (err) {
          this.error(err)
          return
        }
        spinner.succeed(`${fileName} has been created`)
      })
    } else {
      this.log(result)
    }
  }
}
