interface ImportMeta {
    env: {
      [key: string]: string | undefined;
      NODE_ENV: 'development' | 'production' | 'test';
    };
  }