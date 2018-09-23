package external

import (
	"github.com/alkemir/eGSA/gsa"
)

func gsacak(cat []byte) {
	SA := make([]gsa.ResultIndex, len(cat))
	LCP := make([]gsa.ResultIndex, len(cat))
	bucket := make([]gsa.ResultIndex, 256)

	putSubstr0_generalized(cat, SA, bucket)


	induceSAl0_generalized(SA, s, bkt, n, K, false, cs, separator);

  induceSAs0_generalized(SA, s, bkt, n, K, false, cs, separator);

  // insert separator suffixes in their buckets
  // bkt[separator]=1; // gsa-is
  for(i=n-3; i>0; i--)
    if(chr(i)==separator)
      SA[bkt[chr(i)]--]=i;

  // now, all the LMS-substrings are sorted and
  //   stored sparsely in SA.

  // compact all the sorted substrings into
  //   the first n1 items of SA.
  // 2*n1 must be not larger than n.
  uint_t n1=0;
  for(i=0; i<n; i++)
    if((SA[i]>0))
      SA[n1++]=SA[i];

  uint_t *SA1=SA, *s1=SA+m-n1;
  uint_t name_ctr;


  name_ctr=nameSubstr_generalized_LCP(SA,LCP,s,s1,n,m,n1,level,cs,separator);

  // stage 2: solve the reduced problem.
  int_t depth=1;

  // recurse if names are not yet unique.
  if(name_ctr<n1)
    depth += SACA_K((int_t*)s1, SA1,
          n1, 0, m-n1, sizeof(int_t), level+1);
  else // get the suffix array of s1 directly.
    for(i=0; i<n1; i++) SA1[s1[i]]=i;

  // stage 3: induce SA(S) from SA(S1).

  getSAlms(SA, (int_t*)s, s1, n, n1, level, cs);


  uint_t *RA=s1;
  int_t *PLCP=LCP+m-n1;//PHI is stored in PLCP array

  //compute the LCP of consecutive LMS-suffixes
  compute_lcp_phi_sparse((int_t*)s, SA1, RA, LCP, PLCP, n1, cs, separator);

  for(i=0; i<n1; i++) SA[i]=s1[SA[i]];
  for(i=n1; i<n; i++) SA[i]=U_MAX;
  for(i=n1;i<n;i++) LCP[i]=0;


  putSuffix0_generalized_LCP(SA, LCP, s, bkt, n, K, n1, cs, separator);
  induceSAl0_generalized_LCP(SA, LCP, s, bkt, n, K, cs, separator);
  induceSAs0_generalized_LCP(SA, LCP, s, bkt, n, K, cs, separator);

return depth;
}

func putSubstr0_generalized(cat []byte, SA []gsa.ResultIndex, bucket []gsa.ResultIndex) {
	getBuckets_k(cat, bucket) // find the end of each bucket.

	// set each item in SA as empty.
	for i := 0; i < len(SA); i++) {
		SA[i]=0
	}

	// gsa-is
	tmp := bucket[separator]-- // shifts one position left of bkt[separator]

	SA[0]=len(cat)-1 // set the single sentinel LMS-substring.
	SA[tmp] = SA[0]-1// insert the last separator at the end of bkt[separator]

	p:=len(cat)-2
	succ := 0 // s[n-2] must be L-type.


	for(i=len(cat)-2; i>0; i--) {
	  curr := (chr(i-1)<chr(i) || (chr(i-1)==chr(i) && succ_t==1) )?1:0;
	  if(curr ==0 && succ ==1){

		if(chr(i)==separator)
		  SA[++bkt[chr(p)]]=0; // removes LMS-positions that induces separator suffixes

		SA[bkt[chr(i)]--]=i;
		p=i;
	  }
	  succ =curr
	}
  }

}
