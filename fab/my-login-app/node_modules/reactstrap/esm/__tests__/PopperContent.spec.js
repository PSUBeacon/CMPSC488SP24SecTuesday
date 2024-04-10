import React from 'react';
import { Popper } from 'react-popper';
import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';
import { PopperContent } from '..';
describe('PopperContent', function () {
  var element;
  beforeEach(function () {
    element = document.createElement('div');
    element.innerHTML = '<p id="outerTarget">This is the popover <span id="target">target</span>.</p>';
    document.body.appendChild(element);
    jest.useFakeTimers();
  });
  afterEach(function () {
    jest.clearAllTimers();
    document.body.removeChild(element);
    element = null;
  });
  it('should NOT render children when isOpen is false', function () {
    render( /*#__PURE__*/React.createElement(PopperContent, {
      target: "target"
    }, "Yo!"));
    expect(screen.queryByText('Yo!')).not.toBeInTheDocument();
  });
  it('should render children when isOpen is true and container is inline', function () {
    render( /*#__PURE__*/React.createElement(PopperContent, {
      target: "target",
      isOpen: true
    }, "Yo!"));
    expect(screen.queryByText('Yo!')).toBeInTheDocument();
  });
  it('should render children when isOpen is true and container is inline and DOM node passed directly for target', function () {
    var targetElement = element.querySelector('#target');
    render( /*#__PURE__*/React.createElement(PopperContent, {
      target: targetElement,
      isOpen: true,
      container: "inline"
    }, "Yo!"));
    expect(screen.queryByText('Yo!')).toBeInTheDocument();
  });
  it('should render an Arrow in the Popper when isOpen is true and container is inline', function () {
    var _render = render( /*#__PURE__*/React.createElement(PopperContent, {
        target: "target",
        isOpen: true,
        container: "inline",
        arrowClassName: "custom-arrow"
      }, "Yo!")),
      container = _render.container;
    expect(container.querySelector('.arrow.custom-arrow')).toBeInTheDocument();
  });
  it('should NOT render an Arrow in the Popper when "hideArrow" is truthy', function () {
    var _render2 = render( /*#__PURE__*/React.createElement(PopperContent, {
        target: "target",
        isOpen: true,
        container: "inline",
        arrowClassName: "custom-arrow",
        hideArrow: true
      }, "Yo!")),
      container = _render2.container;
    expect(container.querySelector('.arrow.custom-arrow')).not.toBeInTheDocument();
  });
  it('should pass additional classNames to the popper', function () {
    render( /*#__PURE__*/React.createElement(PopperContent, {
      className: "extra",
      target: "target",
      isOpen: true,
      container: "inline",
      "data-testid": "rick"
    }, "Yo!"));
    expect(screen.getByTestId('rick')).toHaveClass('extra');
  });
  it('should allow custom modifiers and even allow overriding of default modifiers', function () {
    render( /*#__PURE__*/React.createElement(PopperContent, {
      className: "extra",
      target: "target",
      isOpen: true,
      container: "inline",
      modifiers: [{
        name: 'offset',
        options: {
          offset: [2, 2]
        }
      }, {
        name: 'preventOverflow',
        options: {
          boundary: 'viewport'
        }
      }]
    }, "Yo!"));
    expect(Popper).toHaveBeenCalledWith(expect.objectContaining({
      modifiers: expect.arrayContaining([expect.objectContaining({
        name: 'offset',
        options: {
          offset: [2, 2]
        }
      }), expect.objectContaining({
        name: 'preventOverflow',
        options: {
          boundary: 'viewport'
        }
      })])
    }), {});
  });
  it('should have data-popper-placement of auto by default', function () {
    var _render3 = render( /*#__PURE__*/React.createElement(PopperContent, {
        target: "target",
        isOpen: true,
        container: "inline"
      }, "Yo!")),
      container = _render3.container;
    expect(container.querySelector('div[data-popper-placement="auto"]')).toBeInTheDocument();
  });
  it('should override data-popper-placement', function () {
    var _render4 = render( /*#__PURE__*/React.createElement(PopperContent, {
        placement: "top",
        target: "target",
        isOpen: true,
        container: "inline"
      }, "Yo!")),
      container = _render4.container;
    expect(container.querySelector('div[data-popper-placement="auto"]')).not.toBeInTheDocument();
    expect(container.querySelector('div[data-popper-placement="top"]')).toBeInTheDocument();
  });
  it('should allow for a placement prefix', function () {
    render( /*#__PURE__*/React.createElement(PopperContent, {
      placementPrefix: "dropdown",
      target: "target",
      isOpen: true,
      container: "inline"
    }, "Yo!"));
    expect(screen.getByText('Yo!')).toHaveClass('dropdown-auto');
  });
  it('should allow for a placement prefix with custom placement', function () {
    var _render5 = render( /*#__PURE__*/React.createElement(PopperContent, {
        placementPrefix: "dropdown",
        placement: "top",
        target: "target",
        isOpen: true,
        container: "inline"
      }, "Yo!")),
      container = _render5.container;
    expect(screen.getByText('Yo!')).toHaveClass('dropdown-auto');
    expect(container.querySelector('div[data-popper-placement="top"]')).toBeInTheDocument();
  });
  it('should render custom tag for the popper', function () {
    render( /*#__PURE__*/React.createElement(PopperContent, {
      tag: "main",
      target: "target",
      isOpen: true,
      container: "inline",
      "data-testid": "morty"
    }, "Yo!"));
    expect(screen.getByTestId('morty').tagName.toLowerCase()).toBe('main');
  });
  it('should allow a function to be used as children', function () {
    var renderChildren = jest.fn();
    render( /*#__PURE__*/React.createElement(PopperContent, {
      target: "target",
      isOpen: true
    }, renderChildren));
    expect(renderChildren).toHaveBeenCalled();
  });
  it('should render children properly when children is a function', function () {
    render( /*#__PURE__*/React.createElement(PopperContent, {
      target: "target",
      isOpen: true
    }, function () {
      return 'Yo!';
    }));
    expect(screen.queryByText('Yo!')).toBeInTheDocument();
  });
});