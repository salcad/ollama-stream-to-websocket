import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import userEvent from '@testing-library/user-event';
import Home from '../app/page';

describe('Home component', () => {
  it('renders the initial UI correctly', () => {
    render(<Home />);

    expect(screen.getByLabelText('Model')).toBeInTheDocument();
    expect(screen.getByLabelText('Prompt')).toBeInTheDocument();
    expect(screen.getByText('Send prompt')).toBeInTheDocument();
  });

  it('updates the model input value when changed', () => {
    render(<Home />);
    const modelInput = screen.getByLabelText('Model');

    expect(modelInput).toHaveValue('llama3.1');
  });

  it('updates the prompt textarea value when changed', () => {
    render(<Home />);
    const promptTextarea = screen.getByLabelText('Prompt');

    expect(promptTextarea).toHaveValue('Why is there something rather than nothing?');
  });

  it('disables the button when loading', async () => {
    render(<Home />);
    const button = screen.getByText('Send prompt');

    global.fetch = jest.fn(() =>
      Promise.resolve({
        text: () => Promise.resolve('response text'),
      } as Response)
    ) as jest.Mock;

    userEvent.click(button);
    await waitFor(() => expect(button).toBeDisabled());
  });
  
});

