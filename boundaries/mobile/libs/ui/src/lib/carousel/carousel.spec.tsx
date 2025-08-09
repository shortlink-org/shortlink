import React from 'react';
import { render } from '@testing-library/react-native';

import Carousel from './carousel';

describe('Carousel', () => {
  it('should render successfully', () => {
    const { root } = render(< Carousel />);
    expect(root).toBeTruthy();
  });
});
