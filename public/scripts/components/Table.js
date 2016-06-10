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

function Table(props) {
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
                <RowContainer />
                </thead>
                <tbody>
                </tbody>
            </table>
        </div>
    )
}

module.exports = Table;