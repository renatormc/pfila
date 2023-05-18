import { Routes, Route } from "react-router-dom";
import { AppProvider } from './AppContext';
import CustomRouter from "~/custom_routes/CustomRoutes"
import customHistory from "~/custom_routes/history";
import ProcessesPage from './pages/ProcessesPage';


function App() {

  return (

    <AppProvider>
      <CustomRouter history={customHistory}>
        <Routes>
          
          <Route path="/" element={<ProcessesPage />} />
        </Routes>
      </CustomRouter>
    </AppProvider>
  )
}

export default App
