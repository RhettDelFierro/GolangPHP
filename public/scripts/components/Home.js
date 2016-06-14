var React = require("react");
var Header = require("./Header");
var PropTypes = React.PropTypes;
var StudentFormContainer = require("../containers/StudentFormContainer");
var TableContainer = require("../containers/TableContainer");


function Home(props) {
    return (
        <div>
            <Header />
            <StudentFormContainer isLoggedIn={props.isLoggedIn} user={props.user}
                                  onHandleUpdateStudent={props.onUpdateStudent}
                                  onStudentSubmit={props.onStudentSubmit}/>
            <TableContainer user={props.user}
                            isLoggedIn={props.isLoggedIn}
                            studentInfo={props.studentInfo}
                            studentsLoaded={props.studentsLoaded}
                            onStudentDelete={props.onStudentDelete}/>
        </div>
    )

}

Home.propTypes = {
    isLoggedIn: PropTypes.bool.isRequired,
    user: PropTypes.string,
    onUpdateStudent: PropTypes.func.isRequired,
    onStudentSubmit: PropTypes.func.isRequired
};
module.exports = Home;