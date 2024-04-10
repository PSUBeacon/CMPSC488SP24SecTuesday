import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import Popover from '../Popover';
describe('Popover', function () {
  var element;
  beforeEach(function () {
    element = document.createElement('div');
    element.setAttribute('id', 'popover-target');
    document.body.appendChild(element);
  });
  afterEach(function () {
    document.body.removeChild(element);
  });
  it('should apply popperClassName to popper component', function () {
    var _screen$queryByText;
    render( /*#__PURE__*/React.createElement(Popover, {
      target: "popover-target",
      popperClassName: "boba-was-here",
      isOpen: true
    }, "Bo-Katan Kryze"));
    expect((_screen$queryByText = screen.queryByText('Bo-Katan Kryze')) === null || _screen$queryByText === void 0 ? void 0 : _screen$queryByText.parentElement).toHaveClass('popover show boba-was-here');
  });
  it('should apply arrowClassName to arrow', function () {
    var _render = render( /*#__PURE__*/React.createElement(Popover, {
        target: "popover-target",
        arrowClassName: "boba-was-here",
        isOpen: true
      }, "Bo-Katan Kryze")),
      debug = _render.debug;
    debug();
    expect(document.querySelector('.arrow')).toHaveClass('boba-was-here');
  });
});