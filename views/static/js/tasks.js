//content
//	tasklist
//		taskbox
//			taskname
//			taskdescription
//			other task stuff

var React = require('react');
var ReactDOM = require('react-dom');

var data = [
  {id: 1, name: "Pete Hunt", desc: "This is one comment"},
  {id: 2, name: "Jordan Walke", desc: "This is *another* comment"}
];

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
      		This is a taskbox.
      	</div>      
    );
  }
});





ReactDOM.render(
<TaskList />,
  document.getElementById('content')
);