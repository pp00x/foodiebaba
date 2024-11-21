import { useEffect, useState, useContext } from 'react';
import { useParams } from 'react-router-dom';
import axios from '../api/axios';
import { AuthContext } from '../context/AuthContext';
import { toast } from 'react-toastify';

function RestaurantDetail() {
  const { id } = useParams();
  const [restaurant, setRestaurant] = useState(null);
  const { user } = useContext(AuthContext);

  useEffect(() => {
    axios
      .get(`/restaurants/${id}`)
      .then((response) => setRestaurant(response.data))
      .catch((error) => console.error(error));
  }, [id]);

  if (!restaurant) return <div>Loading...</div>;

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-2">{restaurant.name}</h1>
      <p className="mb-2">{restaurant.category}</p>
      <p className="mb-4">{restaurant.address}</p>
      <p className="mb-4">{restaurant.description}</p>

      {/* Display photos */}
      {restaurant.photos && restaurant.photos.length > 0 && (
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-4">
          {restaurant.photos.map((photo) => (
            <img key={photo.id} src={photo.url} alt="Restaurant" className="w-full h-48 object-cover" />
          ))}
        </div>
      )}

      {/* Upload photos */}
      {user && (
        <UploadPhotos restaurantId={restaurant.id} />
      )}

      {/* Display reviews */}
      <h2 className="text-2xl font-semibold mb-2">Reviews</h2>
      {restaurant.reviews && restaurant.reviews.length > 0 ? (
        restaurant.reviews.map((review) => (
          <div key={review.id} className="border p-4 rounded mb-2">
            <p className="font-semibold">{review.user.username}</p>
            <p>Rating: {review.rating}/5</p>
            <p>{review.comment}</p>
          </div>
        ))
      ) : (
        <p>No reviews yet.</p>
      )}

      {/* Add a review */}
      {user && <AddReview restaurantId={restaurant.id} />}
    </div>
  );
}

function AddReview({ restaurantId }) {
  const [rating, setRating] = useState(5);
  const [comment, setComment] = useState('');
  const { user } = useContext(AuthContext);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post(
        '/reviews',
        { rating, comment, restaurant_id: restaurantId },
        { headers: { Authorization: `Bearer ${user.token}` } },
      );
      toast.success('Review added');
      window.location.reload();
    } catch (error) {
      if (error.response) {
        toast.error(error.response.data.message || 'Failed to add review');
      } else if (error.request) {
        toast.error('Network error, please try again');
      } else {
        toast.error('An error occurred');
      }
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mt-4">
      <h3 className="text-xl font-semibold mb-2">Add a Review</h3>
      <div className="mb-2">
        <label className="block mb-1">Rating</label>
        <select
          className="border px-3 py-2 w-full"
          value={rating}
          onChange={(e) => setRating(e.target.value)}
          required
        >
          {[1, 2, 3, 4, 5].map((num) => (
            <option value={num} key={num}>
              {num}
            </option>
          ))}
        </select>
      </div>
      <div className="mb-2">
        <label className="block mb-1">Comment</label>
        <textarea
          className="border px-3 py-2 w-full"
          value={comment}
          onChange={(e) => setComment(e.target.value)}
          required
        ></textarea>
      </div>
      <button className="bg-blue-500 text-white px-4 py-2">Submit Review</button>
    </form>
  );
}

function UploadPhotos({ restaurantId }) {
  const [selectedFiles, setSelectedFiles] = useState([]);
  const { user } = useContext(AuthContext);

  const handleFileChange = (e) => {
    setSelectedFiles(e.target.files);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (selectedFiles.length === 0) {
      toast.error('Please select at least one photo');
      return;
    }
    const formData = new FormData();
    for (let i = 0; i < selectedFiles.length; i++) {
      formData.append('photos', selectedFiles[i]);
    }

    try {
      await axios.post(
        `/restaurants/${restaurantId}/photos`,
        formData,
        {
          headers: {
            Authorization: `Bearer ${user.token}`,
            'Content-Type': 'multipart/form-data',
          },
        },
      );
      toast.success('Photos uploaded successfully');
      window.location.reload();
    } catch (error) {
      if (error.response) {
        toast.error(error.response.data.message || 'Failed to upload photos');
      } else if (error.request) {
        toast.error('Network error, please try again');
      } else {
        toast.error('An error occurred');
      }
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mt-4">
      <h3 className="text-xl font-semibold mb-2">Upload Photos</h3>
      <input type="file" multiple onChange={handleFileChange} />
      <button type="submit" className="bg-blue-500 text-white px-4 py-2 mt-2">
        Upload
      </button>
    </form>
  );
}

export default RestaurantDetail;