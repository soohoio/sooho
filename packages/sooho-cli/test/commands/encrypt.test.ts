import {expect, test} from '@oclif/test'

describe('encrypt', () => {
  test
    .stdout()
    .command(['encrypt', 'test/commands/Test.sol'])
    .it('runs extract functions from Test.sol', ctx => {
      const result = ctx.stdout.split('\n')
      expect(result[2]).to.equal('[Function]')
      expect(result[3]).to.equal(
        "functionlength()externalviewsimple(5,'sasd')returns(uint){returntodos.length;}"
      )
      expect(result[4]).to.equal('[Function]')
      expect(result[5]).to.equal(
        "functionaddTodo(bool_isActive,string_text)public{TodoItemmemorytodo=TodoItem(_isActive,_text);todos.push(todo);}"
      )
    })
})
