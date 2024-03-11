/* eslint-disable */
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
};

export type Book = {
  __typename?: 'Book';
  attentionBook?: Maybe<Scalars['Boolean']['output']>;
  author?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  itemUrl?: Maybe<Scalars['String']['output']>;
  janCode: Scalars['String']['output'];
  largeImageUrl?: Maybe<Scalars['String']['output']>;
  mediumImageUrl?: Maybe<Scalars['String']['output']>;
  myBook?: Maybe<Scalars['Boolean']['output']>;
  publisherName?: Maybe<Scalars['String']['output']>;
  title: Scalars['String']['output'];
};

export type DeleteBook = {
  janCode: Scalars['String']['input'];
  userid: Scalars['String']['input'];
};

export type DeleteReadingMemo = {
  janCode: Scalars['String']['input'];
  memoId: Scalars['String']['input'];
};

export type FindBook = {
  janCode: Scalars['String']['input'];
  userid: Scalars['String']['input'];
};

export type GetAllMemo = {
  janCode: Scalars['String']['input'];
};

export type GetMemo = {
  janCode: Scalars['String']['input'];
  memoId: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createAttentionBookData: Book;
  createMyBookData: Book;
  createReadingMemo: ReadingMemo;
  deleteMyBookData: Book;
  deleteReadingMemo: Scalars['Int']['output'];
  updateReadingMemo: ReadingMemo;
};


export type MutationCreateAttentionBookDataArgs = {
  input: NewBook;
};


export type MutationCreateMyBookDataArgs = {
  input: NewBook;
};


export type MutationCreateReadingMemoArgs = {
  input: NewReadingMemo;
};


export type MutationDeleteMyBookDataArgs = {
  input: DeleteBook;
};


export type MutationDeleteReadingMemoArgs = {
  input: DeleteReadingMemo;
};


export type MutationUpdateReadingMemoArgs = {
  input: UpdateReadingMemo;
};

export type NewBook = {
  attentionBook?: InputMaybe<Scalars['Boolean']['input']>;
  author?: InputMaybe<Scalars['String']['input']>;
  itemUrl?: InputMaybe<Scalars['String']['input']>;
  janCode: Scalars['String']['input'];
  largeImageUrl?: InputMaybe<Scalars['String']['input']>;
  mediumImageUrl?: InputMaybe<Scalars['String']['input']>;
  myBook?: InputMaybe<Scalars['Boolean']['input']>;
  publisherName?: InputMaybe<Scalars['String']['input']>;
  title: Scalars['String']['input'];
};

export type NewReadingMemo = {
  author?: InputMaybe<Scalars['String']['input']>;
  content?: InputMaybe<Scalars['String']['input']>;
  janCode?: InputMaybe<Scalars['String']['input']>;
  memoId: Scalars['String']['input'];
  title?: InputMaybe<Scalars['String']['input']>;
  userId?: InputMaybe<Scalars['String']['input']>;
};

export type Query = {
  __typename?: 'Query';
  ReadingMemo: ReadingMemo;
  ReadingMemos: Array<ReadingMemo>;
  findAllMyBook?: Maybe<Array<Book>>;
  findBook?: Maybe<Book>;
};


export type QueryReadingMemoArgs = {
  input: GetMemo;
};


export type QueryReadingMemosArgs = {
  input: GetAllMemo;
};


export type QueryFindBookArgs = {
  input: FindBook;
};

export type ReadingMemo = {
  __typename?: 'ReadingMemo';
  author?: Maybe<Scalars['String']['output']>;
  content?: Maybe<Scalars['String']['output']>;
  janCode: Scalars['String']['output'];
  memoId: Scalars['String']['output'];
  title?: Maybe<Scalars['String']['output']>;
  userId: Scalars['String']['output'];
};

export type UpdateBook = {
  attentionBook?: InputMaybe<Scalars['Boolean']['input']>;
  author?: InputMaybe<Scalars['String']['input']>;
  itemUrl?: InputMaybe<Scalars['String']['input']>;
  janCode: Scalars['String']['input'];
  largeImageUrl?: InputMaybe<Scalars['String']['input']>;
  mediumImageUrl?: InputMaybe<Scalars['String']['input']>;
  myBook?: InputMaybe<Scalars['Boolean']['input']>;
  title: Scalars['String']['input'];
  userid: Scalars['String']['input'];
};

export type UpdateReadingMemo = {
  author?: InputMaybe<Scalars['String']['input']>;
  content?: InputMaybe<Scalars['String']['input']>;
  janCode?: InputMaybe<Scalars['String']['input']>;
  memoId: Scalars['String']['input'];
  title?: InputMaybe<Scalars['String']['input']>;
  userId?: InputMaybe<Scalars['String']['input']>;
};

export type FindAllMyBookQueryVariables = Exact<{ [key: string]: never; }>;


export type FindAllMyBookQuery = { __typename?: 'Query', findAllMyBook?: Array<{ __typename?: 'Book', title: string, author?: string | null, janCode: string, publisherName?: string | null, itemUrl?: string | null, largeImageUrl?: string | null, mediumImageUrl?: string | null, myBook?: boolean | null }> | null };

export type CreateMyBookDataMutationVariables = Exact<{
  input: NewBook;
}>;


export type CreateMyBookDataMutation = { __typename?: 'Mutation', createMyBookData: { __typename?: 'Book', title: string, author?: string | null, janCode: string } };

export type FindAllReadingMemoQueryVariables = Exact<{
  input: GetAllMemo;
}>;


export type FindAllReadingMemoQuery = { __typename?: 'Query', ReadingMemos: Array<{ __typename?: 'ReadingMemo', janCode: string, title?: string | null, content?: string | null }> };

export type CreateReadingMemoMutationVariables = Exact<{
  input: NewReadingMemo;
}>;


export type CreateReadingMemoMutation = { __typename?: 'Mutation', createReadingMemo: { __typename?: 'ReadingMemo', janCode: string, title?: string | null, content?: string | null } };

export type UpdateReadingMemoMutationVariables = Exact<{
  input: UpdateReadingMemo;
}>;


export type UpdateReadingMemoMutation = { __typename?: 'Mutation', updateReadingMemo: { __typename?: 'ReadingMemo', memoId: string, janCode: string, title?: string | null, content?: string | null } };

export type DeleteReadingMemoMutationVariables = Exact<{
  input: DeleteReadingMemo;
}>;


export type DeleteReadingMemoMutation = { __typename?: 'Mutation', deleteReadingMemo: number };


export const FindAllMyBookDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"findAllMyBook"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"findAllMyBook"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"author"}},{"kind":"Field","name":{"kind":"Name","value":"janCode"}},{"kind":"Field","name":{"kind":"Name","value":"publisherName"}},{"kind":"Field","name":{"kind":"Name","value":"itemUrl"}},{"kind":"Field","name":{"kind":"Name","value":"largeImageUrl"}},{"kind":"Field","name":{"kind":"Name","value":"mediumImageUrl"}},{"kind":"Field","name":{"kind":"Name","value":"myBook"}}]}}]}}]} as unknown as DocumentNode<FindAllMyBookQuery, FindAllMyBookQueryVariables>;
export const CreateMyBookDataDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createMyBookData"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewBook"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createMyBookData"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"author"}},{"kind":"Field","name":{"kind":"Name","value":"janCode"}}]}}]}}]} as unknown as DocumentNode<CreateMyBookDataMutation, CreateMyBookDataMutationVariables>;
export const FindAllReadingMemoDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"findAllReadingMemo"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"GetAllMemo"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"ReadingMemos"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"janCode"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"content"}}]}}]}}]} as unknown as DocumentNode<FindAllReadingMemoQuery, FindAllReadingMemoQueryVariables>;
export const CreateReadingMemoDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createReadingMemo"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewReadingMemo"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createReadingMemo"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"janCode"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"content"}}]}}]}}]} as unknown as DocumentNode<CreateReadingMemoMutation, CreateReadingMemoMutationVariables>;
export const UpdateReadingMemoDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updateReadingMemo"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UpdateReadingMemo"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateReadingMemo"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"memoId"}},{"kind":"Field","name":{"kind":"Name","value":"janCode"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"content"}}]}}]}}]} as unknown as DocumentNode<UpdateReadingMemoMutation, UpdateReadingMemoMutationVariables>;
export const DeleteReadingMemoDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"deleteReadingMemo"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"DeleteReadingMemo"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"deleteReadingMemo"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}]}]}}]} as unknown as DocumentNode<DeleteReadingMemoMutation, DeleteReadingMemoMutationVariables>;