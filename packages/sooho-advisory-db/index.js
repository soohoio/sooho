'use strict';
const yaml = require('js-yaml');
const fs   = require('fs');
const path = require('path');

const appDirectory = fs.realpathSync(process.cwd());

const resolveOwn = relativePath => path.resolve(__dirname, relativePath);

function getAdvisoryDB() {
  let data;
   try {
     const filePath = resolveOwn('ethereum/solidity/cve/2018-12056.yml')
     data = yaml.safeLoad(
       fs.readFileSync(filePath, 'utf8')
     );
     return data;
   } catch (e) {
     return new Error(e);
   }
}

module.exports = { getAdvisoryDB };
