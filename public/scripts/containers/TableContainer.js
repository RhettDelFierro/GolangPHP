var React = require("react");

var TableContainer = React.createClass({
    render: function () {
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
                    </thead>
                    <tbody>
                    </tbody>
                </table>
            </div>
        )
    }
});

module.exports = TableContainer;