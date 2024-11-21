import { useContext } from 'react';
import { AuthContext } from '../context/AuthContext';
import { Navigate } from 'react-router-dom';

function PrivateRoute({ children, adminOnly = false }) {
  const { user } = useContext(AuthContext);

  if (!user) {
    // If not logged in, redirect to login page
    return <Navigate to="/login" />;
  }

  if (adminOnly && user.role !== 'admin') {
    // If not an admin, redirect to unauthorized page or home
    return <Navigate to="/" />;
  }

  return children;
}

export default PrivateRoute;