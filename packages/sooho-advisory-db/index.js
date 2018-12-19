'use strict';
const yaml = require('js-yaml');
const fs   = require('fs');
const path = require('path');

const appDirectory = fs.realpathSync(process.cwd());
const cveDirectory = 'ethereum/solidity/cve'
const resolveOwn = relativePath => path.resolve(__dirname, relativePath);

function getAdvisoryDB() {
  let data = [];
   try {
     const cveDir = resolveOwn(cveDirectory)
     fs.readdirSync(cveDir)
       .map(item => resolveOwn(`${cveDirectory}/${item}`))
       .forEach(filePath => data.push(yaml.safeLoad(
         fs.readFileSync(filePath, 'utf8')
       )));
     return data;
   } catch (e) {
     throw new Error(e);
   }
}

module.exports = { getAdvisoryDB };
