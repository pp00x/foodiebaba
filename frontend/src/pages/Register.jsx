import { useNavigate } from 'react-router-dom';
import axios from '../api/axios';
import { Formik, Form, Field, ErrorMessage } from 'formik';
import * as Yup from 'yup';
import { toast } from 'react-toastify';

function Register() {
  const navigate = useNavigate();

  const validationSchema = Yup.object({
    username: Yup.string().required('Required'),
    email: Yup.string().email('Invalid email').required('Required'),
    password: Yup.string().min(6, 'Minimum 6 characters').required('Required'),
  });

  const handleSubmit = async (values, { setSubmitting }) => {
    try {
      await axios.post('/register', values);
      toast.success('Registration successful');
      navigate('/login');
    } catch (error) {
      if (error.response) {
        toast.error(error.response.data.message || 'Registration failed');
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
    <div className="flex items-center justify-center h-screen">
      <Formik
        initialValues={{ username: '', email: '', password: '' }}
        validationSchema={validationSchema}
        onSubmit={handleSubmit}
      >
        {({ isSubmitting }) => (
          <Form className="w-96 p-6 bg-white rounded shadow">
            <h2 className="text-2xl font-semibold mb-4">Register</h2>
            <div className="mb-4">
              <label className="block mb-1">Username</label>
              <Field
                type="text"
                name="username"
                className="w-full border px-3 py-2"
              />
              <ErrorMessage name="username" component="div" className="text-red-500 text-sm" />
            </div>
            <div className="mb-4">
              <label className="block mb-1">Email</label>
              <Field
                type="email"
                name="email"
                className="w-full border px-3 py-2"
              />
              <ErrorMessage name="email" component="div" className="text-red-500 text-sm" />
            </div>
            <div className="mb-4">
              <label className="block mb-1">Password</label>
              <Field
                type="password"
                name="password"
                className="w-full border px-3 py-2"
              />
              <ErrorMessage name="password" component="div" className="text-red-500 text-sm" />
            </div>
            <button
              type="submit"
              className="w-full bg-green-500 text-white py-2"
              disabled={isSubmitting}
            >
              Register
            </button>
          </Form>
        )}
      </Formik>
    </div>
  );
}

export default Register;