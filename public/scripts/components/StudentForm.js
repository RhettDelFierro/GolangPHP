var React = require("react");

function StudentForm(props){
    //remember to add the onClicks.
    console.log("within student form: ", props.isLoggedIn, props.user);
    return (
        <div className="student-add-form col-xs-12 col-md-push-9 col-md-3">
            <h4>Add Student</h4>

            <div className="input-group form-group">
            <span className="input-group-addon">
                <span className="glyphicon glyphicon-user"></span>
            </span>
                <input ref="student" type="text" className="form-control" name="studentName" id="studentName" placeholder="Student Name"/>
            </div>

            <p className="text-danger" id="regex_name"></p>

            <div className="input-group form-group">
            <span className="input-group-addon">
                <span className="glyphicon glyphicon-list-alt"></span>
            </span>
                <input ref="course" type="text" className="form-control" name="course" id="course" placeholder="Student Course"/>
            </div>

            <p className="text-danger" id="regex_course"></p>

            <div className="input-group form-group">
            <span className="input-group-addon">
                <span className="glyphicon glyphicon-education"></span>
            </span>
                <input ref="grade" type="text" className="form-control" name="studentGrade" id="studentGrade" placeholder="Student Grade"/>
            </div>

            <p className="text-danger" id="regex_grade"></p>

            <button type="button" className={props.isLoggedIn === true
            ? "btn btn-success form-group active"
            : "btn btn-success form-group disabled"} >Add</button>
            {' '}
            <button type="button" className="btn btn-success form-group">Load Grades</button>
            {' '}
            <button type="button" className="btn btn-default form-group">Cancel</button>

            <p className="text-danger" id="extra_error"></p>
        </div>
    )
}

module.exports = StudentForm;