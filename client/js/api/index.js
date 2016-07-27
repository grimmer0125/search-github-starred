import repos from './repos';

const api = {
  ...repos,
};

export default api;

// function startSession(token) {
//   return fetch(`${baseUrl}/session`, {
//     method: 'PUT', delete(for delete session) , body is token

// cookie -> token
// function requestToken(credentials) {
//   return fetch(`${api.getBaseApiUrl()}/token/`, {
//     method: 'POST',
//     body: JSON.stringify(credentials), put cookie in body
