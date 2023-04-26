using System;


namespace Maximum
{
    class Program
    {
        static void Main(String[] args)
        {
            string[] s = Console.ReadLine().Split();
            int m = Convert.ToInt32(s[0]);
            int[] ball = new int[m];

            for (int i = 0; i < Convert.ToInt32(s[1]); i++) {
                string[] d = Console.ReadLine().Split();
                int a = Convert.ToInt32(d[0]) - 1;
                int b = Convert.ToInt32(d[1]) - 1; 

                if (i == 0)
                {
                    for (int n = 0; n < m; n++)
                    {
                        ball[n] = n + 1;
                    }
                }

                int c = ball[a];
                ball[a] = ball[b];
                ball[b] = c;

                if (i == Convert.ToInt32(s[1]) - 1)
                {
                    foreach (int num in ball)
                    {
                        Console.Write(num + " ");
                    }
                }
            }
        }
    }
}