import {expect, test} from '@oclif/test'

describe('encrypt', () => {
  test
    .stdout()
    .command(['encrypt', `Todo.sol`])
    .it('runs encrypt with Todo.sol', ctx => {
      expect(ctx.stdout).to.contain('[Function]')
    })
})
