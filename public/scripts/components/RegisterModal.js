var React = require("react");
var Bootstrap = require("react-bootstrap");
var Button = Bootstrap.Button;
var Modal = Bootstrap.Modal;
var Popover = Bootstrap.Popover;
var Tooltip = Bootstrap.Tooltip;
var OverlayTrigger = Bootstrap.OverlayTrigger;
var PropTypes = React.PropTypes;

function RegisterModal(props) {
    return (
        <div>
            <Button
                bsClass="btn btn-primary navbar-btn"
                onClick={props.onOpen}
            >
                Register
            </Button>

            <Modal show={props.modalToggle} onHide={props.onClose}>
                <Modal.Header closeButton>
                    <Modal.Title>Modal heading</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <h4>Text in a modal</h4>
                    <p>Duis mollis, est non commodo luctus, nisi erat porttitor ligula.</p>

                    <hr />

                    <h4>Overflowing text to show scroll behavior</h4>
                    <p>Cras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ac facilisis in,
                        egestas eget quam. Morbi leo risus, porta ac consectetur ac, vestibulum at eros.</p>
                    <p>Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Vivamus sagittis lacus vel
                        augue laoreet rutrum faucibus dolor auctor.</p>
                    <p>Aenean lacinia bibendum nulla sed consectetur. Praesent commodo cursus magna, vel scelerisque
                        nisl consectetur et. Donec sed odio dui. Donec ullamcorper nulla non metus auctor fringilla.</p>
                </Modal.Body>
                <Modal.Footer>
                    <Button onClick={props.onClose}>Close</Button>
                </Modal.Footer>
            </Modal>
        </div>
    )
}

RegisterModal.propTypes = {
    modalToggle: PropTypes.bool.isRequired,
    onOpen: PropTypes.func.isRequired,
    onClose: PropTypes.func.isRequired
};

module.exports = RegisterModal;
