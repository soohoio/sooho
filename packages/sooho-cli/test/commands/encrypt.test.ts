import {expect, test} from '@oclif/test'

describe('encrypt', () => {
  test
    .stdout()
    .command(['encrypt', 'test/commands/Test/Test.sol'])
    .it('accepts file', ctx => {
      const result = JSON.parse(ctx.stdout)

      expect(result).to.deep.equal({
        version: '0.3.0',
        errors: [],
        success: {
          constructors: [],
          functions: [{
            filePath: "test/commands/Test/Test.sol",
            loc: {
              endLine: 20,
              startLine: 17
            },
            signature: "c1804baf32d7f6c23426803ffd8d9456"
          }, {
            filePath: "test/commands/Test/Test.sol",
            loc: {
              endLine: 25,
              startLine: 22
            },
            signature: "5e3c146b68293fe4808f0dd8c1a88c8d"
          }
        ]
      },
      totalFiles: 1,
      totalLines: 26,
      totalSigns: 2
    })
  })

  test
    .stdout()
    .command(['encrypt', 'test/commands/Test'])
    .it('accepts folder', ctx => {
      const result = JSON.parse(ctx.stdout)

      expect(result).to.deep.equal({
        version: '0.3.0',
        errors: [],
        success: {
          constructors: [],
          functions: [{
            filePath: "test/commands/Test/Test.sol",
            loc: {
              endLine: 20,
              startLine: 17
            },
            signature: "c1804baf32d7f6c23426803ffd8d9456"
          }, {
            filePath: "test/commands/Test/Test.sol",
            loc: {
              endLine: 25,
              startLine: 22
            },
            signature: "5e3c146b68293fe4808f0dd8c1a88c8d"
          }
        ]
      },
      totalFiles: 1,
      totalLines: 26,
      totalSigns: 2
    })
  })
})
