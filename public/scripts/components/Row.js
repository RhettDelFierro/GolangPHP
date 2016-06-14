var React = require("react");
var RowContainer = require("../containers/RowContainer");


//these components may need a wrapper. May.
//however,
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
        //delete this.props.key

    },
    render: function () {
        return ( <td>
                <button className="btn btn-danger" onClick={this.deleteRecord}>Delete</button>
            </td>
        )
    }
});

function Record(props) {
    return (
        <tr>
            <Student student={props.student}/>
            <Course course={props.course}/>
            <Grade grade={props.grade}/>
            <Button key={props.key}/>
        </tr>
    )
}

function Row(props) {
    return props.studentsLoaded === true ?
        <Record student={props.student}
                course={props.course}
                grade={props.grade}
                key={props.key}/>
        : <tr></tr>
}

module.exports = Row;