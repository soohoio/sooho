import * as path from 'path'

export const onlySolidity = filePath => path.extname(filePath) === '.sol'
