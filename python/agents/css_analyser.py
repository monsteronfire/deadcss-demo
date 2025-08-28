from typing import TypedDict
from langgraph.graph import StateGraph, Graph, END


class GraphState(TypedDict):
    message: str
    question: str


# Nodes
def welcome_node(state: GraphState) -> dict:
    state["message"] = "Welcome to the test agent"
    return state


def question_node(state: GraphState) -> dict:
    state["question"] = "How's it going?"
    return state


def display_node(state: GraphState) -> dict:
    print(state["message"])
    print(state["question"])
    return {}


def create_graph() -> Graph:
    # Init a graph workflow
    workflow = StateGraph(GraphState)

    # Define nodes to be used to build flow of graph
    workflow.add_node("welcome", welcome_node)
    workflow.add_node("question", question_node)
    workflow.add_node("display", display_node)

    # Create graph connections and flow
    workflow.set_entry_point("welcome")
    workflow.add_edge("welcome", "question")
    workflow.add_edge("question", "display")
    workflow.add_edge("display", END)

    # Return compiled, invokable (runnable) application
    return workflow.compile()


def analyse_css():
    # Set up initial state
    initial_state = {}

    # Invoke the app with an empty initial state (dict)
    flowapp = create_graph()
    flowapp.invoke(initial_state)
