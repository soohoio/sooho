#!/usr/bin/env node

const fs = require("fs");
const path = require("path");
const parser = require("solidity-parser-antlr");
const process = require("process");
const semver = require("semver");
const findNodeModules = require("find-node-modules");

class Graph {
  constructor() {
    this.node = {};
  }

  add(from, to) {
    if (!this.node[from]) {
      this.node[from] = [];
    }

    this.node[from].push(to);
  }
}

function resolvePath(baseDir, filePath) {
  let p = path.join(baseDir, filePath);
  p = path.normalize(p);
  if (fs.existsSync(p)) {
    return p;
  } else {
    const nodeModules = findNodeModules();
    for (let nodeModule of nodeModules) {
      let m = path.join(nodeModule, filePath);
      if (fs.existsSync(m)) {
        return m;
      }
    }
    throw new Error(`module path, ${filePath} is not found`);
  }
}

function getInfo(filePath) {
  let content = fs.readFileSync(filePath, "utf8");

  // get ast with solidity parser
  let ast = parser.parse(content);
  let imports = [];
  let version;
  
  // parse import paths in file
  parser.visit(ast, {
    ImportDirective: function(node) {
      /*
       * path - import path
       * unitAlias / symbolAliases - maybe i can use later
       */
      imports.push(node.path);
    },
    PragmaDirective: function(node) {
      /*
       * value - version sepcified in current file
       */
      version = node.value;
    }
  });

  return {
    imports: imports,
    version: version
  };
}

async function getVersion(versions) {
  let caret = [];
  let pinned = [];
  // unique & group
  [...new Set(versions)].forEach(version => {
    if (version.includes("^")) {
      caret.push(version.substring(1));
    } else {
      pinned.push(version);
    }
  });
  
  let version = "";
  if (caret.length) {
    version = "^" + caret.reduce((a, b) => semver.gt(a, b) ? a : b);
  }

  if (pinned.length) {
    version = pinned.reduce((a, b) => semver.gt(a, b) ? a : b);
  }

  return version;
}

async function flatten(filePath, baseDir, log) {
  let visited = [];
  let order = [];
  let versions = [];
  let graph = new Graph();

  let target = path.resolve(filePath);
  baseDir = path.resolve(baseDir);

  // topological sorting of import graph, cutting cycles
  const toposort = async curr => {
    visited.push(curr);

    let info = getInfo(curr);
    versions.push(info.version);

    for (let dep of info.imports) {
      dep = resolvePath(baseDir, dep);
      graph.add(curr, dep);
      if (!visited.includes(dep)) {
        await toposort(dep);
      }
    }

    order.push(curr);
  };

  await toposort(target);

  let version = await getVersion(versions);

  log("pragma solidity " + version + ";\n");

  const PRAGAMA_SOLIDITY_VERSION_REGEX = /^\s*pragma\ssolidity\s+(.*?)\s*;/m;
  const IMPORT_SOLIDITY_REGEX = /^\s*import(\s+).*$/gm;

  for (let p of order) {
    let content = fs.readFileSync(p, "utf8");
    let pure = content
      .replace(PRAGAMA_SOLIDITY_VERSION_REGEX, "")
      .replace(IMPORT_SOLIDITY_REGEX, "");
    log("\n\n" + pure.trim());
  }
}

async function main(args) {
  try {
    let ret = "";
    await flatten(args[0], args[1], str => (ret += str));
    return ret;
  } catch(err) {
    console.error(err);
  }
}

if (require.main == module) {
  main(process.argv.slice(2));
}

module.exports = async function(filePath, baseDir) {
  try {
    let ret = "";
    await flatten(filePath, baseDir, str => (ret += str));
    return ret;
  } catch(err) {
    console.error(err);
  }
};
