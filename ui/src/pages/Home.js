import {
  List,
  ListItem,
  ListIcon,
  Button,
  Input,
  Alert,
  AlertIcon,
  Highlight,
} from "@chakra-ui/react";
import React, { useState, useEffect } from "react";
import axios from "axios";
import {
  InfoOutlineIcon,
  SmallCloseIcon,
  CheckCircleIcon,
  AddIcon,
} from "@chakra-ui/icons";

const Home = () => {
  const [jsondata, setjsondata] = useState([]);
  const [todo, setTodo] = useState("");

  const getdata = async () => {
    const url = "http://localhost:8080/";

    try {
      const response = await axios.get(url);
      setjsondata(response.data.data); // Fetch and update todo list
    } catch (error) {
      console.error("Error fetching json: ", error.message);
    }
  };

  const handleaddtodo = async () => {
    try {
      const response = await axios.post("http://localhost:8080/create", {
        name: todo,
      });
      if (response.status === 201) {
        setTodo("");
        <Alert status="success">
          <AlertIcon />
          Todo added. Fire on!
        </Alert>;
        await getdata();
        alert("Failed to add todo. Status code: " + response.status);
      }
    } catch (error) {
      alert("An error occurred: " + error.message);
    }
  };

  const handleDeleteTodo = async (id) => {
    const url = `http://localhost:8080/delete/${id}`;
    try {
      const response = await axios.delete(url);
      if (response.status === 200) {
        <Alert status="success">
          <AlertIcon />
          Todo deleted!
        </Alert>;
        await getdata(); // Refresh the todo list after deletion
      }
    } catch (error) {
      console.error("Error deleting todo: ", error.message);
    }
  };

  useEffect(() => {
    getdata();
  }, []);

  return (
    <div className="my-2">
      <div className="w-2/4 m-auto flex flex-row items-center justify-between">
        <div className="w-3/4">
          <Input
            type="text"
            name="item"
            placeholder="Todo name"
            value={todo}
            onChange={(e) => setTodo(e.target.value)}
            size="sm"
            variant="flushed"
          />
        </div>

        <div className="w-1/6 flex items-end justify-end">
          <Button
            colorScheme="blue"
            onClick={handleaddtodo}
            size={"sm"}
            leftIcon={<AddIcon />}
          >
            Add
          </Button>
        </div>
      </div>

      <div className="w-2/4 m-auto mt-2">
        <List spacing={3}>
          {jsondata.length > 0 ? (
            jsondata.map((item) => (
              <div
                key={item.ID}
                className="border rounded flex flex-row items-center justify-between m-auto p-2"
              >
                <div className="w-1/4">
                  <ListItem>
                    <ListIcon as={InfoOutlineIcon} color="green.500" />
                    {item.Name || "Unnamed"}
                  </ListItem>
                </div>

                <div className="flex w-1/4 justify-evenly">
                  <Button
                    colorScheme="green"
                    onClick={() => handleDeleteTodo(item.ID)}
                    leftIcon={<CheckCircleIcon />}
                    size={"xs"}
                  >
                    Complete
                  </Button>

                  <Button
                    colorScheme="red"
                    onClick={() => handleDeleteTodo(item.ID)}
                    leftIcon={<SmallCloseIcon />}
                    size={"xs"}
                    variant="outline"
                  >
                    Delete
                  </Button>
                </div>
              </div>
            ))
          ) : (
            <ListItem>
              <Highlight
                query="spotlight"
                styles={{ px: "1", py: "1", bg: "orange.100" }}
              >
                No todo's found
              </Highlight>
            </ListItem>
          )}
        </List>
      </div>
    </div>
  );
};

export default Home;
