import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Time: { input: any; output: any; }
};

export type Authentication = {
  __typename?: 'Authentication';
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  identifier: Scalars['String']['output'];
  provider: AuthenticationProvider;
  updatedAt: Scalars['Time']['output'];
  user: User;
  userID: Scalars['ID']['output'];
};

export enum AuthenticationProvider {
  Facebook = 'facebook',
  Google = 'google'
}

export type Query = {
  __typename?: 'Query';
  user: User;
};


export type QueryUserArgs = {
  id: Scalars['ID']['input'];
  withOptions: UserWithOptions;
};

export type Task = {
  __typename?: 'Task';
  createdAt: Scalars['Time']['output'];
  cron: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  isActive: Scalars['Boolean']['output'];
  nextExecutionTime: Scalars['Time']['output'];
  taskHistories?: Maybe<Array<TaskHistory>>;
  tradingAccount: TradingAccount;
  tradingAccountID: Scalars['ID']['output'];
  type: Scalars['String']['output'];
  updatedAt: Scalars['Time']['output'];
};

export type TaskHistory = {
  __typename?: 'TaskHistory';
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  isSuccess: Scalars['Boolean']['output'];
  task: Task;
  taskID: Scalars['ID']['output'];
  updatedAt: Scalars['Time']['output'];
};

export type TradingAccount = {
  __typename?: 'TradingAccount';
  createdAt: Scalars['Time']['output'];
  currency: Scalars['String']['output'];
  exchange: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  identifier: Scalars['String']['output'];
  ip: Scalars['String']['output'];
  tasks?: Maybe<Array<Task>>;
  updatedAt: Scalars['Time']['output'];
  user: User;
  userID: Scalars['ID']['output'];
};

export type User = {
  __typename?: 'User';
  authentications?: Maybe<Array<Authentication>>;
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  timezone: Scalars['String']['output'];
  tradingAccounts?: Maybe<Array<TradingAccount>>;
  updatedAt: Scalars['Time']['output'];
};

export type UserWithOptions = {
  withAuthentications: Scalars['Boolean']['input'];
  withTradingAccounts: Scalars['Boolean']['input'];
};

export type GetUserQueryVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type GetUserQuery = { __typename?: 'Query', user: { __typename?: 'User', id: string, name: string } };


export const GetUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"user"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}},{"kind":"Argument","name":{"kind":"Name","value":"withOptions"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"withTradingAccounts"},"value":{"kind":"BooleanValue","value":false}},{"kind":"ObjectField","name":{"kind":"Name","value":"withAuthentications"},"value":{"kind":"BooleanValue","value":false}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<GetUserQuery, GetUserQueryVariables>;