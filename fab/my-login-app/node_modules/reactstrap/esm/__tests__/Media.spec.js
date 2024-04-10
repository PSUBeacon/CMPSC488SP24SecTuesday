import React from 'react';
import { render, screen } from '@testing-library/react';
import { Media } from '..';
import { testForChildrenInComponent, testForCustomTag, testForDefaultClass, testForDefaultTag } from '../testUtils';
describe('Media', function () {
  it('should render a div tag by default', function () {
    testForDefaultTag(Media, 'div');
  });
  it('should render an h4 tag by default for heading', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      "data-testid": "test",
      heading: true
    }));
    expect(screen.getByTestId('test').tagName.toLowerCase()).toBe('h4');
  });
  it('should render an a tag by default Media with an href', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      href: "#",
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test').tagName.toLowerCase()).toBe('a');
  });
  it('should render an img tag by default for object', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      object: true,
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test').tagName.toLowerCase()).toBe('img');
  });
  it('should render an img tag by default Media with a src', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      src: "#",
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test').tagName.toLowerCase()).toBe('img');
  });
  it('should render a ul tag by default for list', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      list: true,
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test').tagName.toLowerCase()).toBe('ul');
  });
  it('should pass additional classNames', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      className: "extra",
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test')).toHaveClass('extra');
  });
  it('should render custom tag', function () {
    testForCustomTag(Media);
  });
  it('should render body', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      body: true,
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test')).toHaveClass('media-body');
  });
  it('should render heading', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      heading: true,
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test')).toHaveClass('media-heading');
  });
  it('should render left', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      left: true,
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test')).toHaveClass('media-left');
  });
  it('should render right', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      right: true,
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test')).toHaveClass('media-right');
  });
  it('should render top', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      top: true,
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test')).toHaveClass('media-top');
  });
  it('should render bottom', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      bottom: true,
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test')).toHaveClass('media-bottom');
  });
  it('should render middle', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      middle: true,
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test')).toHaveClass('media-middle');
  });
  it('should render object', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      object: true,
      "data-testid": "test"
    }));
    expect(screen.getByTestId('test')).toHaveClass('media-object');
  });
  it('should render media', function () {
    testForDefaultClass(Media, 'media');
  });
  it('should render list', function () {
    render( /*#__PURE__*/React.createElement(Media, {
      list: true,
      "data-testid": "test"
    }, /*#__PURE__*/React.createElement(Media, {
      tag: "li"
    }), /*#__PURE__*/React.createElement(Media, {
      tag: "li"
    }), /*#__PURE__*/React.createElement(Media, {
      tag: "li"
    })));
    expect(screen.getByTestId('test').querySelectorAll('li').length).toBe(3);
  });
  it('should render children', function () {
    testForChildrenInComponent(Media);
  });
});