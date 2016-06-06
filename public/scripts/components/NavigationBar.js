var React = require("react");
var PropTypes = React.PropTypes;
var Bootstrap = require("react-bootstrap");
var Navbar = Bootstrap.Navbar;
var Nav = Bootstrap.Nav;
var FormGroup = Bootstrap.FormGroup;
var FormControl = Bootstrap.FormControl;
var Button = Bootstrap.Button;
var NavItem = Bootstrap.NavItem;
var MenuItem = Bootstrap.MenuItem;
var NavDropdown = Bootstrap.NavDropdown;


function LoggedIn(props) {
    //probably just use a dropdown instead of just text.
    return (
        <Navbar.Collapse bsStyles="success">
            <Navbar.Text pullRight>
                Signed in as:
            </Navbar.Text>

            <NavDropdown eventKey={3} title={props.user} id="basic-nav-dropdown">
                <MenuItem divider/>
                <MenuItem eventKey={3.1}>Logout</MenuItem>
            </NavDropdown>
        </Navbar.Collapse>
    )
}

function NotLoggedIn(props) {
    //remember, two different FormContainers.
    return (
        <Navbar.Collapse>
            <Nav>
                <Navbar.Form pullRight>
                    <FormGroup>
                        <FormControl type="text" placeholder="Search"/>
                    </FormGroup>
                    {' '}
                    <Button type="submit">Submit</Button>
                </Navbar.Form>
            </Nav>
            <Nav pullRight>
                <NavItem eventKey={1} href="#">Link Right</NavItem>
                <NavItem eventKey={2} href="#">Link Right</NavItem>
            </Nav>
        </Navbar.Collapse>
    )
}

function NavigationBar(props) {
    return (
        <Navbar staticTop fluid>
            <Navbar.Header>
                <Navbar.Brand>
                    <a href="#">Student Grade Table</a>
                </Navbar.Brand>
                <Navbar.Toggle />
            </Navbar.Header>

            {props.isLoggedIn === true
                ? <LoggedIn user={props.user}/>
                : <NotLoggedIn />
            }

        </Navbar>
    )
}

NavigationBar.propTypes = {
    isLoggedIn: PropTypes.bool.isRequired,
    user: PropTypes.string
};

module.exports = NavigationBar;