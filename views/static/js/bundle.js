(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);var f=new Error("Cannot find module '"+o+"'");throw f.code="MODULE_NOT_FOUND",f}var l=n[o]={exports:{}};t[o][0].call(l.exports,function(e){var n=t[o][1][e];return s(n?n:e)},l,l.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({1:[function(require,module,exports){
//content
//	tasklist
//		taskbox
//			taskname
//			taskdescription
//			other task stuff

var TaskList = React.createClass({
		displayName: 'TaskList',
		render: function () {
				return React.createElement('div', { className: "tasklist" }, "!!HOW DO WE USE SOME OTHER VARIABLE????!!");
		}
});

var TaskBox = React.createClass({ displayName: 'TaskBox',
		render: function () {
				return React.createElement('div', { className: "commentBox" }, "Hello, world! I am a CommentBox.");
		}
});

ReactDOM.render(React.createElement(TaskBox, null), document.getElementById('content'));

},{}]},{},[1]);
