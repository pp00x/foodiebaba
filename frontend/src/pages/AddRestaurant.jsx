import { useContext } from 'react';
import axios from '../api/axios';
import { AuthContext } from '../context/AuthContext';
import { Formik, Form, Field, ErrorMessage } from 'formik';
import * as Yup from 'yup';
import { toast } from 'react-toastify';

function AddRestaurant() {
  const { user } = useContext(AuthContext);

  const validationSchema = Yup.object({
    name: Yup.string().required('Required'),
    category: Yup.string().required('Required'),
    address: Yup.string().required('Required'),
    description: Yup.string().required('Required'),
  });

  const handleSubmit = async (values, { setSubmitting, resetForm }) => {
    try {

       // Destructure required fields
      const { name, category, address, description } = values;

          // Create payload with only required fields
      const payload = { name, category, address, description };
      await axios.post(
        '/restaurants',
        payload,
        { headers: { Authorization: `Bearer ${user.token}` } },
      );
      toast.success('Restaurant added and pending approval');
      resetForm();
    } catch (error) {
      if (error.response) {
        toast.error(error.response.data.message || 'Failed to add restaurant');
      } else if (error.request) {
        toast.error('Network error, please try again');
      } else {
        toast.error('An error occurred');
      }
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-4">Add a New Restaurant</h1>
      <Formik
        initialValues={{ name: '', category: '', address: '', description: '' }}
        validationSchema={validationSchema}
        onSubmit={handleSubmit}
      >
        {({ isSubmitting }) => (
          <Form className="max-w-lg">
            <div className="mb-2">
              <label className="block mb-1">Name</label>
              <Field
                name="name"
                className="border px-3 py-2 w-full"
                required
              />
              <ErrorMessage name="name" component="div" className="text-red-500 text-sm" />
            </div>
            <div className="mb-2">
              <label className="block mb-1">Category</label>
              <Field
                name="category"
                className="border px-3 py-2 w-full"
                required
              />
              <ErrorMessage name="category" component="div" className="text-red-500 text-sm" />
            </div>
            <div className="mb-2">
              <label className="block mb-1">Address</label>
              <Field
                name="address"
                className="border px-3 py-2 w-full"
                required
              />
              <ErrorMessage name="address" component="div" className="text-red-500 text-sm" />
            </div>
            <div className="mb-2">
              <label className="block mb-1">Description</label>
              <Field
                name="description"
                as="textarea"
                className="border px-3 py-2 w-full"
                required
              />
              <ErrorMessage name="description" component="div" className="text-red-500 text-sm" />
            </div>
            <button
              type="submit"
              className="bg-green-500 text-white px-4 py-2"
              disabled={isSubmitting}
            >
              Submit
            </button>
          </Form>
        )}
      </Formik>
    </div>
  );
}

export default AddRestaurant;