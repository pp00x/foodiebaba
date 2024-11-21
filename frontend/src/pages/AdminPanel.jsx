import { useEffect, useState, useContext } from 'react';
import axios from '../api/axios';
import { AuthContext } from '../context/AuthContext';
import { toast } from 'react-toastify';

function AdminPanel() {
  const [pendingRestaurants, setPendingRestaurants] = useState([]);
  const { user } = useContext(AuthContext);

  useEffect(() => {
    axios
      .get('/admin/restaurants/pending', { headers: { Authorization: `Bearer ${user.token}` } })
      .then((response) => setPendingRestaurants(response.data))
      .catch((error) => console.error(error));
  }, [user]);

  const handleApprove = async (id) => {
    try {
      await axios.put(`/admin/restaurants/${id}/approve`, null, {
        headers: { Authorization: `Bearer ${user.token}` },
      });
      setPendingRestaurants(pendingRestaurants.filter((r) => r.id !== id));
      toast.success('Restaurant approved');
    } catch (error) {
      if (error.response) {
        toast.error(error.response.data.message || 'Failed to approve restaurant');
      } else if (error.request) {
        toast.error('Network error, please try again');
      } else {
        toast.error('An error occurred');
      }
    }
  };

  const handleReject = async (id) => {
    try {
      await axios.put(`/admin/restaurants/${id}/reject`, null, {
        headers: { Authorization: `Bearer ${user.token}` },
      });
      setPendingRestaurants(pendingRestaurants.filter((r) => r.id !== id));
      toast.success('Restaurant rejected');
    } catch (error) {
      if (error.response) {
        toast.error(error.response.data.message || 'Failed to reject restaurant');
      } else if (error.request) {
        toast.error('Network error, please try again');
      } else {
        toast.error('An error occurred');
      }
    }
  };

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-4">Admin Panel - Pending Restaurants</h1>
      {pendingRestaurants.length > 0 ? (
        pendingRestaurants.map((restaurant) => (
          <div key={restaurant.id} className="border p-4 rounded mb-2">
            <h2 className="text-xl font-semibold">{restaurant.name}</h2>
            <p>{restaurant.category}</p>
            <p>{restaurant.address}</p>
            <p>{restaurant.description}</p>
            <div className="mt-2">
              <button
                className="bg-green-500 text-white px-3 py-1 mr-2"
                onClick={() => handleApprove(restaurant.id)}
              >
                Approve
              </button>
              <button
                className="bg-red-500 text-white px-3 py-1"
                onClick={() => handleReject(restaurant.id)}
              >
                Reject
              </button>
            </div>
          </div>
        ))
      ) : (
        <p>No pending restaurants.</p>
      )}
    </div>
  );
}

export default AdminPanel;