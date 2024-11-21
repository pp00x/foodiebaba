import { useEffect, useState } from 'react';
import axios from '../api/axios';
import { Link } from 'react-router-dom';
import { toast } from 'react-toastify';

function Restaurants() {
  const [restaurants, setRestaurants] = useState([]);
  const [page, setPage] = useState(1);
  const [limit] = useState(10); // Items per page
  const [name, setName] = useState('');
  const [category, setCategory] = useState('');
  const [totalPages, setTotalPages] = useState(1);

  useEffect(() => {
    axios
      .get('/restaurants', {
        params: {
          page,
          limit,
          name,
          category,
        },
      })
      .then((response) => {
        console.log('API Response:', response.data);
        setRestaurants(response.data); // Adjusted here
        // If totalPages is provided differently, adjust accordingly
        // For example, setTotalPages(response.data.totalPages);
      })
      .catch((error) => {
        console.error(error);
        toast.error('Failed to fetch restaurants');
      });
  }, [page, name, category]);

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-4">Restaurants</h1>

      <div className="flex mb-4">
        <input
          type="text"
          placeholder="Search by name"
          className="border px-3 py-2 mr-2"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <input
          type="text"
          placeholder="Filter by category"
          className="border px-3 py-2"
          value={category}
          onChange={(e) => setCategory(e.target.value)}
        />
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        {restaurants.map((restaurant) => (
          <Link
            to={`/restaurants/${restaurant.id}`}
            key={restaurant.id}
            className="border p-4 rounded hover:shadow"
          >
            <h2 className="text-xl font-semibold">{restaurant.name}</h2>
            <p>{restaurant.category}</p>
            <p>{restaurant.address}</p>
          </Link>
        ))}
      </div>

      {/* Pagination controls can be adjusted or removed based on API capabilities */}
      <div className="flex justify-center mt-4">
        <button
          className="px-3 py-1 border mr-2"
          onClick={() => setPage(page - 1)}
          disabled={page === 1}
        >
          Previous
        </button>
        <span className="px-3 py-1">
          Page {page}
          {totalPages && ` of ${totalPages}`}
        </span>
        <button
          className="px-3 py-1 border ml-2"
          onClick={() => setPage(page + 1)}
          disabled={totalPages ? page === totalPages : false}
        >
          Next
        </button>
      </div>
    </div>
  );
}

export default Restaurants;