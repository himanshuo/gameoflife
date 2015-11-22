//content
//	tasklist
//		taskbox
//			taskname
//			taskdescription
//			other task stuff

var TaskList = React.createClass({
	displayName: 'TaskList',
	render: function(){
		return (
			React.createElement(
				'div', 
				{className: "tasklist"},
				"!!HOW DO WE USE SOME OTHER VARIABLE??????!!"
			)
		)
	}
});

var TaskBox = React.createClass({displayName: 'TaskBox',
  render: function(){
    return (
      React.createElement('div', {className: "commentBox"},
        "Hello, world! I am a CommentBox."
      )
    );
  }
});


ReactDOM.render(
  React.createElement(TaskBox, null),
  document.getElementById('content')
);