/**
 * Created by Rhett on 6/9/16.
 */

//function to render the table

//function to render the row

//function to render each td

//function to render the delete button

//most will have a lot of the same styling, maybe use a wrapper.

var React = require("react");
var RowContainer = require("../containers/RowContainer");
var userFunctions = require("../utils/userFunctions");

function Course(props) {
    return <td>{props.course}</td>
}

function Grade(props) {
    return <td>{props.grade}</td>
}

function Student(props) {
    return <td>{props.student}</td>
}

var Button = React.createClass({
    deleteRecord: function () {
        console.log("delete started");
        userFunctions.deleteStudent(this.props.id)
            .then(function(data){
                this.props.onStudentDelete(data.student.ID)
            }.bind(this))
    },
    render: function () {
        return (
            <td>
                <button className="btn btn-danger" onClick={this.deleteRecord}>Delete</button>
            </td>
        )
    }
});

function Row(props) {
    return (
        <tr>
            <Student student={props.studentInfo.name}/>
            <Course course={props.studentInfo.course}/>
            <Grade grade={props.studentInfo.grade}/>
            <Button id={props.studentInfo.id} onStudentDelete={props.onStudentDelete}/>
        </tr>
    )
}
//{props.studentsLoaded === true
//    ? {rows}
//    : <tr></tr>}

function Table(props) {
    var rows = props.studentInfo.map(function (studentData, index) {
        return <Row user={props.user} studentInfo={studentData}
                    onStudentDelete={props.onStudentDelete} index={index} key={studentData.id}/>;
    });
    return (
        <div className="student-list-container col-xs-12 col-md-9 col-md-pull-3">
            <table className="student-list table">
                <thead>
                <tr>
                    <th>Student Name</th>
                    <th>Student Course</th>
                    <th>Student Grade</th>
                    <th>Operations</th>
                </tr>
                {rows}
                </thead>
                <tbody>
                </tbody>
            </table>
        </div>
    )
}

module.exports = Table;