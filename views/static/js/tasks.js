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
	getInitialState: function(){
		return {data:[]};
	},
	getAllTasks: function(){
		$.ajax({
			url: this.props.url,
			dataType:'json',
			cache:false,
			success: function(data){
				this.setState({data:data});
			}.bind(this),
			error: function(xhr, status, err){
				console.error(this.props.url, status, err.toString());
			}.bind(this)
		});
	},
	componentDidMount: function(){
		 this.getAllTasks();
    		setInterval(this.getAllTasks, this.props.pollInterval);
	},
	render: function(){
		var taskBoxes = this.state.data.map(function(task){
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
      	<div className="taskbox" >
      		<div className="id">
      			{this.props.id}
      		</div>
      		<div className="name">
      			<a  href="http://www.google.com">{this.props.name}</a>

      		</div>
      	</div>      
    );
  }
});


var TaskCreationForm = React.createClass({
  getInitialState: function(){
  	return {name:""};
  },
  handleNameChange: function(e){
  	this.setState({name: e.target.value});
  },
  createTask: function(name){	
  	//todo: add sprintf
  	$.ajax({
		type:"PUT",
		url: this.props.url,
		contentType: "application/x-www-form-urlencoded; charset=UTF-8",
		data: "name="+name, // todo: how to properly append
		cache:false,
		success: function(data){
			var newTask = JSON.parse(data);
			console.log("new task created: "+newTask.name+" :"+newTask.id);
		}.bind(this),
		error: function(xhr, status, err){
			console.error(this.props.url, status, err.toString());
		}.bind(this)
		});
  },
  handleSubmit: function(e){
  	e.preventDefault();
  	var name = this.state.name.trim();
  	if(!name){
  		return;
  	}
  	this.createTask(name);
  	this.setState({name: ""});
  },
  render: function() {
    return (
      <form className="taskCreationForm" onSubmit={this.handleSubmit}>
        <input type="text" placeholder="Task Name"  value={this.state.name} onChange={this.handleNameChange} />
        <input type="submit" value="Post" />
      </form>
    );
  }
});

var TaskSection = React.createClass({displayName:"TaskSection",
	render: function(){
		return(
			<div className="taskSection">
				<TaskList url="/task/" pollInterval={2000} />	
				<TaskCreationForm url="/task/" />
			</div>
		);
	}
});

ReactDOM.render(
<TaskSection />,
  document.getElementById('content')
);