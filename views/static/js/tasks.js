//content
//	tasklist
//		taskbox
//			taskname
//			taskdescription
//			other task stuff

var React = require('react');
var ReactDOM = require('react-dom');

data = [{"id":0,"name":"hi"},{"id":0,"name":"my"},{"id":0,"name":"nm"},{"id":0,"name":"asfd"},{"id":0,"name":"asdf"},{"id":0,"name":"fd"},{"id":0,"name":"asdfadf"},{"id":0,"name":"l"},{"id":0,"name":"asdf"},{"id":0,"name":"asdfadsfLAST"}];

var TaskList = React.createClass({
	displayName: 'TaskList',
	render: function(){
		var taskBoxes = this.props.data.map(function(task){
			return (
				<TaskBox id={task.id} name={task.name} />
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
      		<div className="id">
      			{this.props.id}
      		</div>
      		<div className="name">
      			{this.props.name}

      		</div>
      	</div>      
    );
  }
});



//"localhost:8080/task/"

ReactDOM.render(
<TaskList data={data} />,
  document.getElementById('content')
);