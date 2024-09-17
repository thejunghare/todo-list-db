import {
  Route,
  createBrowserRouter,
  createRoutesFromElements,
  RouterProvider,
} from "react-router-dom";
import RootLayout from "./layouts/RootLayout";
import NotFound from "./pages/NotFound";
import TodoForm from "./pages/TodoForm";
import Home from "./pages/Home";
import { ChakraProvider } from "@chakra-ui/react";
import theme from "./theme";

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route element={<RootLayout />}>
      <Route index element={<Home />} />
      <Route path="create" element={<TodoForm />} />

      <Route path="*" element={<NotFound />} />
    </Route>
  )
);

function App() {
  return (
    <ChakraProvider theme={theme} className="p-3">
      <RouterProvider router={router} />
    </ChakraProvider>
  );
}

export default App;
