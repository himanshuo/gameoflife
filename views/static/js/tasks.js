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
		var taskBoxes = this.props.data.map(function(task){
			return (
				<TaskBox name={task.name} desc={task.desc} />
			);
		});
		return (
			<div className="tasklist">
			{taskBoxes}
			</div>	
		);
	}
});



var TaskBox = React.createClass({displayName: 'TaskBox',
  render: function(){
    return (
      	<div className="taskbox">
      		<div className="name">
      			{this.props.name}
      		</div>
      		<div className="desc">
      			{this.props.desc}

      		</div>
      	</div>      
    );
  }
});





ReactDOM.render(
<TaskList data={data} />,
  document.getElementById('content')
);