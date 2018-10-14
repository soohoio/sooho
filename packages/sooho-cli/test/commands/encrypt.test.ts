import {expect, test} from '@oclif/test'

describe('encrypt', () => {
  test
    .stdout()
    .command(['encrypt', 'test/commands/Test.sol'])
    .it('runs extract functions from Test.sol without abstraction', ctx => {
      const result = ctx.stdout.split('\n')

      expect(result).to.deep.equal([
        'Extracting signatures in test/commands/Test.sol without abstraction',
        '[Function]',
        JSON.stringify({ startLine: 17, endLine: 20 }),
        'c1804baf32d7f6c23426803ffd8d9456',
        '[Function]',
        JSON.stringify({ startLine: 22, endLine: 25 }),
        '5e3c146b68293fe4808f0dd8c1a88c8d',
        ''
      ])
    })

  test
    .stdout()
    .command(['encrypt', 'test/commands/Test.sol', '-a'])
    .it('runs extract functions from Test.sol with abstraction', ctx => {
      const result = ctx.stdout.split('\n')

      expect(result).to.deep.equal([
        'Extracting signatures in test/commands/Test.sol with abstraction',
        '[Function]',
        JSON.stringify({ startLine: 17, endLine: 20 }),
        '2ca728bd15bd3a7616d189e355fe6431',
        '[Function]',
        JSON.stringify({ startLine: 22, endLine: 25 }),
        '5e3c146b68293fe4808f0dd8c1a88c8d',
        ''
      ])
    })
})
