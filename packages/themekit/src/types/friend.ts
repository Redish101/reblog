export type Friend = {
  name: string;
  avatar: string;
  url: string;
  desc: string;
};

export type FriendList = {
  count: number;
  friends: Friend[];
};
