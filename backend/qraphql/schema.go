package graphql

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	# The query type, represents all of the entry points into our object graph
	type Query {
	  users(): [User]!
	  user(id: ID!): User
	}
	type Mutation {
	  createUser(firstName: String!, lastName: String!, email: String!, password: String!, nickname: String!): User
	  updateUser(id: ID!, firstName: String!, lastName: String!, email: String!, password: String!, nickname: String!): User
	  deleteUser(id: ID!): ID
	}
	type User {
	  id: ID!
	  firstName: String!
	  lastName: String!
	  email: String!
	  password: String!
      nickname: String!
	}
	type Survey {
	  id: ID!
	  title: String!
	  description: String!
	  schedule: Schedule!
	  questions: [Question]!
	  recipients: String!
	}
	type Question {
	  id:ID!
	  title: String!
	  details: String!
	  qtype: QType!
	  category: String!
	}
	type QType {
	  id: ID!
	  title: String!
	  type: String!
	}
	type Schedule{
	  id: ID!
	  value: String!
	}
`
