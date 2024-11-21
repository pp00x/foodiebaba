import { useContext } from 'react';
import { AuthContext } from '../context/AuthContext';
import axios from '../api/axios';
import { useNavigate } from 'react-router-dom';
import { Formik, Form, Field, ErrorMessage } from 'formik';
import * as Yup from 'yup';
import { toast } from 'react-toastify';

function Login() {
  const { login } = useContext(AuthContext);
  const navigate = useNavigate();

  const validationSchema = Yup.object({
    email: Yup.string().email('Invalid email').required('Required'),
    password: Yup.string().required('Required'),
  });

  const handleSubmit = async (values, { setSubmitting }) => {
    try {
      const response = await axios.post('/login', values);
      login(response.data);
      toast.success('Login successful');
      navigate('/');
    } catch (error) {
      if (error.response) {
        toast.error(error.response.data.message || 'Login failed');
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
        initialValues={{ email: '', password: '' }}
        validationSchema={validationSchema}
        onSubmit={handleSubmit}
      >
        {({ isSubmitting }) => (
          <Form className="w-96 p-6 bg-white rounded shadow">
            <h2 className="text-2xl font-semibold mb-4">Login</h2>
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
              className="w-full bg-blue-500 text-white py-2"
              disabled={isSubmitting}
            >
              Login
            </button>
          </Form>
        )}
      </Formik>
    </div>
  );
}

export default Login;