import { Link } from 'react-router-dom';
import { useContext } from 'react';
import { AuthContext } from '../context/AuthContext';

function Navbar() {
  const { user, logout } = useContext(AuthContext);

  return (
    <nav className="bg-gray-800 text-white p-4 flex justify-between">
      <Link to="/" className="font-bold text-xl">FoodieBaba</Link>
      <div>
        {user ? (
          <>
            <span className="mr-4">Hello, {user.username}</span>
            {user.role === 'admin' && (
              <Link to="/admin" className="mr-4">Admin Panel</Link>
            )}
            <Link to="/add-restaurant" className="mr-4">Add Restaurant</Link>
            <button onClick={logout} className="bg-red-500 px-3 py-1 rounded">Logout</button>
          </>
        ) : (
          <>
            <Link to="/login" className="mr-4">Login</Link>
            <Link to="/register" className="mr-4">Register</Link>
          </>
        )}
      </div>
    </nav>
  );
}

export default Navbar;