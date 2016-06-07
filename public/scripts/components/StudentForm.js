var React = require("react");

function StudentForm(props){
    //remember to add the onClicks.
    return (
        <div className="student-add-form col-xs-12 col-md-push-9 col-md-3">
            <h4>Add Student</h4>

            <div className="input-group form-group">
            <span className="input-group-addon">
                <span className="glyphicon glyphicon-user"></span>
            </span>
                <input type="text" className="form-control" name="studentName" id="studentName" placeholder="Student Name"/>
            </div>

            <p className="text-danger" id="regex_name"></p>

            <div className="input-group form-group">
            <span className="input-group-addon">
                <span className="glyphicon glyphicon-list-alt"></span>
            </span>
                <input type="text" className="form-control" name="course" id="course" placeholder="Student Course"/>
            </div>

            <p className="text-danger" id="regex_course"></p>

            <div className="input-group form-group">
            <span className="input-group-addon">
                <span className="glyphicon glyphicon-education"></span>
            </span>
                <input type="text" className="form-control" name="studentGrade" id="studentGrade" placeholder="Student Grade"/>
            </div>

            <p className="text-danger" id="regex_grade"></p>

            <button type="button" className="btn btn-success form-group">Add</button>
            {' '}
            <button type="button" className="btn btn-success form-group">Load Grades</button>
            {' '}
            <button type="button" className="btn btn-default form-group">Cancel</button>

            <p className="text-danger" id="extra_error"></p>
        </div>
    )
}

module.exports = StudentForm;