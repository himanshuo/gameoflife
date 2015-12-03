//content
//	tasklist
//		taskbox
//			taskname
//			taskdescription
//			other task stuff

var React = require('react');
var ReactDOM = require('react-dom');

var TaskList = React.createClass({
	displayName: 'TaskList',
	render: function(){
		return (
			<div className="tasklist">
				<TaskBox />
				<TaskBox />
				<TaskBox />
				<TaskBox />
			</div>	
		);
	}
});



var TaskBox = React.createClass({displayName: 'TaskBox',
  render: function(){
    return (
      	<div className="taskbox">
      		Contents of TaskBox
      	</div>      
    );
  }
});





ReactDOM.render(
<TaskList />,
  document.getElementById('content')
);