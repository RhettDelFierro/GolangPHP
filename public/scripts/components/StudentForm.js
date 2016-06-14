var React = require("react");

function StudentForm(props) {
    //remember to add the onClicks.
    return (
        <form className="student-add-form col-xs-12 col-md-push-9 col-md-3" onSubmit={props.onSubmitStudent}>
            <h4>Add Student</h4>

            <div className="input-group form-group">
            <span className="input-group-addon">
                <span className="glyphicon glyphicon-user"></span>
            </span>
                <input onChange={props.onUpdateStudent} type="text" className="form-control" name="studentName"
                       id="studentName"
                       placeholder="Student Name" value={props.student}/>
            </div>

            <p className="text-danger" id="regex_name"></p>

            <div className="input-group form-group">
            <span className="input-group-addon">
                <span className="glyphicon glyphicon-list-alt"></span>
            </span>
                <input onChange={props.onUpdateCourse} type="text" className="form-control" name="course" id="course"
                       placeholder="Student Course" value={props.course}/>
            </div>

            <p className="text-danger" id="regex_course"></p>

            <div className="input-group form-group">
            <span className="input-group-addon">
                <span className="glyphicon glyphicon-education"></span>
            </span>
                <input onChange={props.onUpdateGrade} type="text" className="form-control" name="studentGrade"
                       id="studentGrade"
                       placeholder="Student Grade" value={props.grade}/>
            </div>

            <p className="text-danger" id="regex_grade"></p>

            <button type="submit" className={props.isLoggedIn === true
            ? "btn btn-success form-group active"
            : "btn btn-success form-group disabled"}>Add
            </button>
            {' '}
            <button type="button" className="btn btn-success form-group" onClick={props.onPopulate}>Load Grades</button>
            {' '}
            <button type="button" className="btn btn-default form-group">Cancel</button>

            <p className="text-danger" id="extra_error"></p>
        </form>
    )

};

module.exports = StudentForm;