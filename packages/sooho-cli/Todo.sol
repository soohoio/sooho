pragma solidity ^0.4.21;

contract Todo {
  struct TodoItem {
    bool isActive;
    string text;
  }

  event LogHello();

  modifier simple(uint test, string tmp) {
    _;
  }

  TodoItem[] public todos;

  function length() external view simple(5, 'sasd') returns (uint) {
    emit LogHello();
    return todos.length;
  }

  function addTodo(bool _isActive, string _text) public {
    TodoItem memory todo = TodoItem(_isActive, _text);
    todos.push(todo);
  }
}
